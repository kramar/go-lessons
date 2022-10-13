package pangram

import "strings"

// IsPangram checks if given string is a pangram
func IsPangram(s string) bool {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	counter := 0

	for _, char := range alphabet {
		if strings.Index(strings.ToLower(s), string(char)) != -1 {
			counter++
		}
	}

	return counter == len(alphabet)
}
