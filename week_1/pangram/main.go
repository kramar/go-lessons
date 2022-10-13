package pangram

import "strings"

// IsPangram checks if given string is a pangram
func IsPangram(s string) bool {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, char := range alphabet {
		if strings.Index(strings.ToLower(s), string(char)) == -1 {
			return false
		}
	}

	return true
}
