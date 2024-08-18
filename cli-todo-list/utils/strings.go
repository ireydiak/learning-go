package utils

import "bytes"

func ReplaceUntilChar(source string, replacement string, stopToken rune) string {
	var buf bytes.Buffer
	found := false
	for _, char := range source {
		if char == stopToken {
			found = true
		}
		if found {
			buf.WriteRune(char)
		}
	}
	return replacement + buf.String()
}
