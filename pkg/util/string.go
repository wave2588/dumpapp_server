package util

import "unicode/utf8"

func StringCount(content string) int {
	return utf8.RuneCountInString(content)
}
