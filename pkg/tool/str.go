package tool

import "strings"

func HandlePostgreArray(value string) []string {
	if len(value) <= 2 {
		return nil
	}

	return strings.Split(value[1:len(value)-1], ",")
}
