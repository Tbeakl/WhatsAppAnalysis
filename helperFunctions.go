package main

import (
	"strings"
	"time"
)

//Copied from https://stackoverflow.com/questions/50982524/how-to-gracefully-iterate-a-date-range-in-go
func rangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

func mostCommonString(values ...string) string {
	stringCounts := make(map[string]int)
	for _, str := range values {
		stringCounts[str]++
	}
	mostCommonString := ""
	mostCommonCount := -1
	for str, count := range stringCounts {
		if count > mostCommonCount {
			mostCommonString = str
			mostCommonCount = count
		}
	}
	return mostCommonString
}

func averageStringLength(values ...string) float64 {
	totalLength := 0
	for _, str := range values {
		totalLength += len(str)
	}
	return float64(totalLength) / float64(len(values))
}

func averageNumberOfWords(values ...string) float64 {
	totalLength := 0
	for _, str := range values {
		totalLength += len(strings.Split(str, " "))
	}
	return float64(totalLength) / float64(len(values))
}

func findMaxStringLengthWords(values ...string) int {
	maxLength := 0
	for _, element := range values {
		if len(strings.Split(element, " ")) > maxLength {
			maxLength = len(strings.Split(element, " "))
		}
	}
	return maxLength
}

func findMaxStringLengthCharacters(values ...string) int {
	maxLength := 0
	for _, element := range values {
		if len(element) > maxLength {
			maxLength = len(element)
		}
	}
	return maxLength
}

func extractUserMessageContent(messages []userMessage) []string {
	output := make([]string, len(messages))
	for i, message := range messages {
		output[i] = message.Content
	}
	return output
}

func extractDateMessageContent(messages []dateMessage) []string {
	output := make([]string, len(messages))
	for i, message := range messages {
		output[i] = message.Content
	}
	return output
}

func extractMessageContent(messages []message) []string {
	output := make([]string, len(messages))
	for i, message := range messages {
		output[i] = message.Content
	}
	return output
}