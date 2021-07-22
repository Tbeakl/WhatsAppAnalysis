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

func findMaxMessageLengthWords(messages []message) int {
	maxLength := 0
	for _, element := range messages {
		if len(strings.Split(element.Content, " ")) > maxLength {
			maxLength = len(strings.Split(element.Content, " "))
		}
	}
	return maxLength
}

func findMaxMessageLengthCharacters(messages []message) int {
	maxLength := 0
	for _, element := range messages {
		if len(element.Content) > maxLength {
			maxLength = len(element.Content)
		}
	}
	return maxLength
}
