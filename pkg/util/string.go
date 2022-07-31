package util

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
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

func CheckURLValid(URL string) error {
	if 0 == len(URL) {
		return errors.New("url is empty")
	}
	parse, err := url.Parse(URL)
	if err != nil {
		return err
	}

	if "http" != parse.Scheme && "https" != parse.Scheme {
		return errors.New("url Scheme illegal. err:" + parse.Scheme)
	}
	re := regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)+([A-Za-z]|[A-Za-z][A-Za-z0-9\-]*[A-Za-z0-9])$`)

	u := strings.Split(parse.Host, ":")
	if 2 == len(u) && 0 != len(u[1]) {
		result := re.FindAllStringSubmatch(u[0], -1)
		if result == nil {
			return errors.New("url illegal. err:" + u[0])
		}
		return nil
	}

	result := re.FindAllStringSubmatch(parse.Host, -1)
	if result == nil {
		return errors.New("url illegal. err:" + parse.Host)
	}

	return nil
}
