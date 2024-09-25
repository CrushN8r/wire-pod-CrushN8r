package wirepod_ttr

import (
	"regexp"
	"time"

	"github.com/sashabaranov/go-openai"
)

const (
	ActionSayText        = iota  // 0
	ActionPlayAnimation          // 1
	ActionPlayAnimationWI        // 2
	ActionGetImage	             // 3
	ActionNewRequest             // 4
	ActionPlaySound              // 5
)

type RobotAction struct {
	Action    int
	Parameter string
}

type LLMCommand struct {
	Command         string
	Description     string
	ParamChoices    string
	Action          int
	SupportedModels []string
}

// create function which parses from LLM and makes a struct of RobotActions
var (
	playAnimationWIDescription = "Enhances storytelling by playing an animation on the robot without interrupting speech. This should be used frequently to animate responses and engage the audience. Choose from parameters like happy, veryHappy, sad, verySad, angry, frustrated, dartingEyes, confused, thinking, celebrate, or love to complement the dialogue and maintain context."

	playAnimationDescription = "Interrupts speech to play an animation on the robot. Only use this when directed explicitly to play an animation, as it halts ongoing speech. Parameters include happy, veryHappy, sad, verySad, angry, frustrated, dartingEyes, confused, thinking, celebrate, or love for expressing emotions and reactions."

	getImageDescription = "Retrieves an image from the robot's camera and displays it in the next message. Use this command to conclude a sentence or response when prompted by the user or when describing visual content, such as what the robot sees, with options for front or lookingUp perspectives. If asked 'What do you see in front of you?' or similar, default to taking a photo. Inform the user of your action before using the command."

	newVoiceRequestDescription = "Starts a new voice command from the robot. Use this if you want more input from the user/if you want to carry out a conversation. You are the only one who can end it in this case. This goes at the end of your response, if you use it."

//	var playSoundDescription = "Plays a sound on the robot."

	LLMCommandsParamChoices = "happy, veryHappy, sad, verySad, angry, frustrated, dartingEyes, confused, thinking, celebrate, love"
)

var ValidLLMCommands []LLMCommand = []LLMCommand{
	{
		Command:         "playAnimationWI",
		Description:     playAnimationWIDescription,
		ParamChoices:    LLMCommandsParamChoices,
		Action:          ActionPlayAnimationWI,
		SupportedModels: []string{"all"},
	},
	{
		Command:         "playAnimation",
		Description:     playAnimationDescription,
		ParamChoices:    LLMCommandsParamChoices,
		Action:          ActionPlayAnimation,
		SupportedModels: []string{"all"},
	},
	{
		Command:         "getImage",
		Description:     getImageDescription,
		ParamChoices:    "front, looking-Up, ahead-Of-You, in-Front-Of-You, Above-You",
		Action:          ActionGetImage,
		SupportedModels: []string{openai.GPT4o, openai.GPT4oMini},
	},
	{
		Command:         "newVoiceRequest",
		Description:     newVoiceRequestDescription,
		ParamChoices:    "now",
		Action:          ActionNewRequest,
		SupportedModels: []string{"all"},
	},
	/*{
	 	Command:      "playSound",
	 	Description:  playSoundDescription,
	 	ParamChoices: "drumroll",
	 	Action:       ActionPlaySound,
	},*/	
}


// Declare variables to store the last processed input and its processing time at the package level
var (
    lastProcessedInput string
    lastProcessedTime  time.Time // Track last processed time

    // Combined Precompiled regex patterns
    bracketDecimalPattern       = regexp.MustCompile(`(\s*[\[\{])\.`)
    numeric0To9Pattern          = regexp.MustCompile(`[0-9]`)
    recurringPattern            = regexp.MustCompile(`point\s+(\d+)\s+recurring`)
    spaceDecimalPattern         = regexp.MustCompile(`(\d+)\s+([./])`)
    unwantedSpacePattern        = regexp.MustCompile(`\.(\s+)`)
    fractionPattern             = regexp.MustCompile(`^\s*(\d+)\s*/\s*(\d+)\s*$`) // Pattern for fractions

    closingBracketPattern       = regexp.MustCompile(`\\\)`)
    closingBracketSquarePattern  = regexp.MustCompile(`\\\]`)
    latexDelimiterPattern       = regexp.MustCompile(`\\\(`)
    latexFractionRegexp         = regexp.MustCompile(`\\(?:frac|dfrac)\{([^}]+)\}\{([^}]+)\}`)
    numericPattern              = regexp.MustCompile(`^\d+(\.\d+)?$`)
    openingBracketPattern       = regexp.MustCompile(`\\\[`)

    safeCharacterPattern        = regexp.MustCompile(`[^\w\s\.{}\^\[\]\/]`)
    simpleFractionRegexp        = regexp.MustCompile(`(\d+)\s*/\s*(\d+)`)
)
