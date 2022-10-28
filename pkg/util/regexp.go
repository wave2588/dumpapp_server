package util

import "regexp"

func RegexpNamedMatch(re *regexp.Regexp, text string) (map[string]string, bool) {
	matches := re.FindStringSubmatch(text)
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" && len(matches) >= i {
			result[name] = matches[i]
		}
	}
	return result, len(matches) > 0
}
