package wirepod_ttr

import (
	"context"
	"errors"
	"io"
	"strings"

	"github.com/fforchino/vector-go-sdk/pkg/vector"
	"github.com/kercre123/wire-pod/chipper/pkg/logger"
	"github.com/kercre123/wire-pod/chipper/pkg/vars"
	"github.com/sashabaranov/go-openai"
)

func SplitLLMResponse(fullRespText string, fullRespSlice []string) (string, []string) {
	// Define the possible separators.
	punctuation := []string{"...", ".'", ".\"", ".", "?", "!"}

	// Loop through the punctuation options to find the first match.
	var sepStr string
	for _, sep := range punctuation {
		if strings.Contains(fullRespText, sep) {
			sepStr = sep
			break
		}
	}

	// If a separator was found, split the text.
	if sepStr != "" {
		splitResp := strings.SplitN(strings.TrimSpace(fullRespText), sepStr, 2)
		if len(splitResp) == 2 {
			fullRespSlice = append(fullRespSlice, strings.TrimSpace(splitResp[0])+sepStr)
			fullRespText = splitResp[1]
		}
	}

	return fullRespText, fullRespSlice
}

func HandleLLMStream(c *openai.Client, ctx context.Context, aireq openai.ChatCompletionRequest, msgs []openai.ChatCompletionMessage, robot *vector.Vector, stopStop chan bool, isDone *bool, fullRespSlice *[]string, tempRespText *string, speakReady chan string) error {
	stream, err := c.CreateChatCompletionStream(ctx, aireq)
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") && vars.APIConfig.Knowledge.Provider == "openai" {
			logger.Println("GPT-4 model cannot be accessed with this API key. You likely need to add more than $5 dollars of funds to your OpenAI account.")
			logger.LogUI("GPT-4 model cannot be accessed with this API key. You likely need to add more than $5 dollars of funds to your OpenAI account.")
			aireq.Model = openai.GPT3Dot5Turbo
			logger.Println("Falling back to " + aireq.Model)
			logger.LogUI("Falling back to " + aireq.Model)
			stream, err = c.CreateChatCompletionStream(ctx, aireq)
			if err != nil {
				logger.Println("OpenAI still not returning a response even after falling back. Erroring.")
				return err
			}
		} else {
			logger.Println("LLM error: " + err.Error())
			return err
		}
	}

	go func() {
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				*isDone = true
				newStr := strings.Join(*fullRespSlice, " ")
				if strings.TrimSpace(newStr) != strings.TrimSpace(*tempRespText) {
					logger.Println("LLM debug: there is content after the last punctuation mark")
					extraBit := strings.TrimPrefix(*tempRespText, newStr)
					*fullRespSlice = append(*fullRespSlice, extraBit)
				}
				if vars.APIConfig.Knowledge.SaveChat {
					Remember(msgs[len(msgs)-1],
						openai.ChatCompletionMessage{
							Role:    openai.ChatMessageRoleAssistant,
							Content: newStr,
						},
						robot.Cfg.SerialNo)
				}
				logger.LogUI("LLM response for " + robot.Cfg.SerialNo + ": " + undoTmpDecPnt(newStr))
				logger.Println("LLM stream finished")
				return
			}

			if err != nil {
				logger.Println("Stream error: " + err.Error())
				return
			}

			*tempRespText += response.Choices[0].Delta.Content
			// Handle decimals
			*tempRespText = doTmpDecPnt(*tempRespText)

			// Split LLM response
			*tempRespText, *fullRespSlice = SplitLLMResponse(*tempRespText, *fullRespSlice)

			if len(*fullRespSlice) > 0 {
				affectedText := strings.TrimSpace((*fullRespSlice)[len(*fullRespSlice)-1]) // Get the last added part
				select {
				case speakReady <- undoTmpDecPnt(affectedText):
				default:
				}
			}
		}
	}()

	return nil
}
