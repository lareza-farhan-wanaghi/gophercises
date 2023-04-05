package hr1

import (
	"regexp"
	"strings"
)

// Camelcase returns the number of English words within a camelcase word
func Camelcase(text string) int32 {
	camelcaseRe := regexp.MustCompile(`^[a-z].*`)
	if !camelcaseRe.MatchString(text) {
		return 0
	}

	var result int32 = 1
	uppercaseRe := regexp.MustCompile(`[A-Z]`)
	allUppercases := uppercaseRe.FindAllString(text, -1)
	if allUppercases != nil {
		result += int32(len(allUppercases))
	}
	return result
}

// CaesarChiper returns the encripted version of the message for the given offset
func CaesarChiper(text string, offsite int32) string {
	aByte := 'a'
	zByte := 'z'
	AByte := 'A'
	ZByte := 'Z'

	var sb strings.Builder
	for _, r := range text {
		if r >= aByte && r <= zByte {
			diff := r - aByte
			sb.WriteRune(aByte + (diff+offsite)%26)
		} else if r >= AByte && r <= ZByte {
			diff := r - AByte
			sb.WriteRune(AByte + (diff+offsite)%26)
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
