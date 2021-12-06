package helper

import (
	"strconv"
	"strings"
)

func InputToIntSlice(input, seperator string) ([]int, error) {
	var i []int
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		for _, substring := range strings.Split(line, seperator) {
			lineInt, err := strconv.Atoi(substring)
			if err != nil {
				return nil, err
			}
			i = append(i, lineInt)
		}
	}
	return i, nil
}

func InputToStringSlice(input string) ([]string, error) {
	var s []string
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		s = append(s, line)
	}
	return s, nil
}
