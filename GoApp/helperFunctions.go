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

//Returns the string which is most common
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

//Calculates the average string length in characters
func averageStringLength(values ...string) float64 {
	totalLength := 0
	for _, str := range values {
		totalLength += len(str)
	}
	return float64(totalLength) / float64(len(values))
}

//Calculates the average number of words in a string
func averageNumberOfWords(values ...string) float64 {
	totalLength := 0
	for _, str := range values {
		totalLength += len(removeEmptyStrings(strings.Split(str, " ")))
	}
	return float64(totalLength) / float64(len(values))
}

func findMaxStringLengthWords(values ...string) int {
	maxLength := 0
	for _, element := range values {
		if len(strings.Split(element, " ")) > maxLength {
			maxLength = len(removeEmptyStrings(strings.Split(element, " ")))
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

func countOfString(element string, values []string) int {
	count := 0
	for _, v := range values {
		if v == element {
			count++
		}
	}
	return count
}

func elementWithLargestSlice(input map[string][]userMessage) string {
	largestLength := 0
	var largestElement string
	for k, v := range input {
		if len(v) > largestLength {
			largestElement = k
			largestLength = len(v)
		}
	}
	return largestElement
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

//From https://gist.github.com/johnpili/84c3064d30a9b041c87e43ba4bcb63a2
// removeEmptyStrings - Use this to remove empty string values inside an array.
// This happens when allocation is bigger and empty
func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func removeMessages(content []string, messagesToRemove ...string) []string {
	var output []string
	for _, v := range content {
		if !contains(messagesToRemove, v) {
			output = append(output, v)
		}
	}
	return output
}

func contains(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func leastSquares(series map[point]float64) line {
	xMean := 0.0
	yMean := 0.0
	totalFrequency := 0.0
	for point, frequency := range series {
		totalFrequency += frequency
		xMean += point.X * frequency
		yMean += point.Y * frequency
	}

	xMean /= totalFrequency
	yMean /= totalFrequency
	covariance := 0.0
	xVariance := 0.0
	for point, frequency := range series {
		xDiff := point.X - xMean
		yDiff := point.Y - yMean
		covariance += xDiff * yDiff * frequency
		xVariance += xDiff * xDiff * frequency
	}
	gradient := covariance / xVariance
	yIntercept := yMean - xMean*gradient
	return line{Gradient: gradient, YIntercept: yIntercept}
}
