package wirepod_ttr

import (
    "strings"
)

var (
    tmpDecPntSpc = "tmpDecPnt " // tmpDecPnt + space
    dotSpc       = ". " // dot + space
)

// before llm response is split
func doTmpDecPnt(rawStr string) string {
    output1 := strings.ReplaceAll(rawStr, ".", "tmpDecPnt") // Replace dot with temporary placeholder
    output2 := strings.ReplaceAll(output1, tmpDecPntSpc, dotSpc) // Replace placeholder+space point with actual dot+space
    return output2
}

// after llm response is split
func undoTmpDecPnt(rawStr string) string {
    output1 := strings.ReplaceAll(rawStr, "tmpDecPnt", ".")
    return output1
}

