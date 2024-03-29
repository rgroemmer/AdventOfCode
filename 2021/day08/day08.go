package main

import (
	"adventOfCode/helper"
	"fmt"
	"os"
	"strings"
)

type Display struct {
	digits []string
	output []string
}

func parseInput(input []string) (displays []Display) {
	for _, line := range input {
		digits := strings.Split(strings.Split(line, "|")[0], " ")
		output := strings.Split(strings.Split(line, "|")[1], " ")
		displays = append(displays, Display{
			digits: digits,
			output: output,
		})
	}
	return
}

func part1(input []string) (result int) {
	displays := parseInput(input)
	for _, display := range displays {
		for _, s := range display.output {
			switch len(s) {
			case 2:
				result++
			case 4:
				result++
			case 3:
				result++
			case 7:
				result++
			}
		}
	}
	return
}

func part2(input []string) (result int) {
	displays := parseInput(input)

	for _, display := range displays {

		digitCodes := make([]string, 10)
		segmentMap := make(map[string]string, 7)

		for _, digit := range display.digits {
			switch len(digit) {
			case 2:
				digitCodes[1] = digit
				display.removeDigit(digit)
			case 4:
				digitCodes[4] = digit
				display.removeDigit(digit)
			case 3:
				digitCodes[7] = digit
				display.removeDigit(digit)
			case 7:
				digitCodes[8] = digit
				display.removeDigit(digit)
			}
		}

		segmentMap["a"] = strings.Trim(digitCodes[7], digitCodes[1])
		segmentMap["g"] = findMissingSegmentCode(display, &digitCodes[9], 6, digitCodes[4]+segmentMap["a"])
		segmentMap["d"] = findMissingSegmentCode(display, &digitCodes[3], 5, digitCodes[1]+segmentMap["a"]+segmentMap["g"])
		segmentMap["b"] = strings.Trim(digitCodes[9], digitCodes[3])
		segmentMap["f"] = findMissingSegmentCode(display, &digitCodes[5], 5, segmentMap["a"]+segmentMap["d"]+segmentMap["g"]+segmentMap["b"])
		segmentMap["e"] = findMissingSegmentCode(display, &digitCodes[6], 6, digitCodes[5])
		segmentMap["c"] = findMissingSegmentCode(display, &digitCodes[2], 5, segmentMap["a"]+segmentMap["d"]+segmentMap["g"]+segmentMap["e"])
		digitCodes[0] = segmentMap["a"] + segmentMap["b"] + segmentMap["c"] + segmentMap["e"] + segmentMap["f"] + segmentMap["g"]

		result += getOutputValue(display.output, digitCodes)
	}

	return
}

func (d *Display) removeDigit(s string) {
	for i := 0; i < len(d.digits); i++ {
		if d.digits[i] == s {
			d.digits[i] = ""
		}
	}
}

func getOutputValue(output []string, codes []string) (result int) {
	for _, digit := range output {
		for i, code := range codes {
			if containsAllCharsAndEqualLen(code, digit) {
				result *= 10
				result += i
				break
			}
		}
	}
	return result
}

func findMissingSegmentCode(display Display, digitCodeResult *string, length int, stringFilter string) string {
	for _, digit := range display.digits {
		if len(digit) == length && containsChars(digit, stringFilter) {
			*digitCodeResult = digit
			display.removeDigit(digit)
		}
	}
	return strings.Trim(*digitCodeResult, stringFilter)
}

func containsAllCharsAndEqualLen(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for _, char := range b {
		if !strings.ContainsRune(a, char) {
			return false
		}
	}
	return true
}

func containsChars(s, chars string) bool {
	for _, c := range chars {
		if !strings.Contains(s, string(c)) {
			return false
		}
	}
	return true
}

func main() {
	r, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input, err := helper.InputToStringSlice(string(r))
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
