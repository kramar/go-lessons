// This is a "stub" file. Це є маленький запуск на вашій згоді.
// It's not a complete solution though; You have to write some code.

// Package leap should have package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear повинен мати коментар documenting it.
func IsLeapYear(year int) bool {
	switch {
	case year % 100 == 0 && year % 400 != 0 :
		return false
	case year % 400 == 0 || year % 4 == 0 :
		return true
	default :
		return false
	}
}
