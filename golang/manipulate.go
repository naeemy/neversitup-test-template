package main

import (
	"math"
	"strconv"
	"strings"
)

func Manipulate(text string) []string {
	if len(text) == 1 {
		return []string{text}
	} else {
		return shuffle(text)
	}
}

func dedup(text string) string {
	m := make(map[string]bool)

	for _, ch := range text {
		m[string(ch)] = true
	}

	result := []string{}

	for k, _ := range m {
		result = append(result, k)
	}

	return strings.Join(result, "")
}

func shuffle(text string) []string {
	m := make(map[string]bool)
	result := []string{}

	for number := range int(math.Pow(float64(len(text)), float64(len(text)))) {
		strNum := strconv.FormatInt(int64(number), len(text))

		for len(strNum) < len(text) {

			strNum = "0" + strNum
		}

		if len(dedup(strNum)) == len(text) {
			composedStr := []string{}
			for _, ch := range strNum {
				i, _ := strconv.Atoi(string(ch))
				composedStr = append(composedStr, string(text[int(i)]))
			}

			m[strings.Join(composedStr, "")] = true
		}

	}

	for key := range m {
		result = append(result, key)
	}

	return result
}
