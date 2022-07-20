package util

import (
	"unicode/utf8"
)

func StringCount(content string) int {
	return utf8.RuneCountInString(content)
}

func CheckUDIDValid(udid string) bool {
	udidLen := StringCount(udid)
	if udidLen != 25 && udidLen != 40 {
		return false
	}
	return true
}
