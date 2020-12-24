package rules

import "strings"

func checkKeyValueString(str string) bool  {
	split := strings.Split(strings.TrimSuffix(str, "\n"), "=")
	if len(split) == 2 {
		return true
	}
	return false
}

func hasComment(str string) bool  {
	split := strings.Split(strings.TrimSuffix(str, "\n"), "#")
	if len(split) > 1 {
		return true
	}
	return false
}