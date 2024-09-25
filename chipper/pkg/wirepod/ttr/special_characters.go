package wirepod_ttr

import (
    "regexp"
    "strings"
    "unicode"

    "golang.org/x/text/transform"
    "golang.org/x/text/unicode/norm"
    "github.com/kercre123/wire-pod/chipper/pkg/logger"
)

// isMn checks if a rune is a non-spacing mark.
func isMn(r rune) bool {
    return unicode.Is(unicode.Mn, r)
}

// normalizeText normalizes the input text by removing non-spacing marks.
func normalizeText(str string) (string, error) {
    t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
    normalizedStr, _, err := transform.String(t, str)
    if err != nil {
        logger.Println("normalizeText error:", err)
        return str, err // Return original input on error
    }
    return normalizedStr, nil
}

// Precompile emoji regex pattern once
var emojiRegex = regexp.MustCompile(emojiPattern)

func removeEmojis(str string) string {
    // Check if emojis exist and clean them
    if emojiRegex.MatchString(str) {
        logger.Println("Emojis detected, cleaning the input.")
        // Clean emojis but retain the original string if the resulting string is empty
        cleanedStr := emojiRegex.ReplaceAllString(str, "")
        return strings.TrimSpace(cleanedStr) // Ensure no leading/trailing spaces
    }
    return str
}

// replaceWords case-insensitively replaces words/phrases in the input string.
func replaceWords(str string) string {
    for find, replace := range phoneticReplacements {
        regex := regexp.MustCompile(`\b(?i)` + regexp.QuoteMeta(find) + `\b`)
        str = regex.ReplaceAllString(str, replace) // Replace words/phrases to allow multiple replacements
    }
    return str
}

// removeSpecialCharacters removes special characters from the input string and provides a user-friendly output if empty.
func removeSpecialCharacters(rawStr string) (string, error) {

    // Normalize the input and log changes; ensure this does not strip punctuation.
    normalizedStr, err := normalizeText(rawStr)
    if err != nil {
        logger.Println("preCleanSpecialCharacters Normalization error:", err)
        return "[No content]", err // Return error as well for better handling
    }

    wordsStr := replaceWords(normalizedStr)

    cleanStr1 := specialCharactersReplacements.Replace(wordsStr)
    cleanStr2 := removeEmojis(cleanStr1) // Clean emojis after initial cleaning

    if strings.TrimSpace(cleanStr2) != "" {
        return strings.TrimSpace(cleanStr2), nil // Return cleaned input
    } else {
        return "", nil // Default return for empty input
    }
}