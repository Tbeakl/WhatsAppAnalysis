package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

// func numberOfMessagesByUsers(messagesByUsers map[string][]userMessage, baseFileName string) {
// 	var numberOfMessages []barChartOutput
// 	for user, messages := range messagesByUsers {
// 		numberOfMessages = append(numberOfMessages, barChartOutput{X: user, Y: len(messages)})
// 	}
// 	jsonData, _ := json.Marshal(numberOfMessages)
// 	ioutil.WriteFile(baseFileName+"\\NumberOfMessages_Users.json", jsonData, os.ModePerm)
// 	var frequencies []int
// 	for _, v := range numberOfMessages {
// 		frequencies = append(frequencies, v.Y)
// 	}
// 	createZipfLawPlot(frequencies, baseFileName+"\\NumberOfMessages_Users_")
// }

func createZipfLawPlot(frequencies []int, baseFileName string) {
	type outputStyle struct {
		Line   line    `json:"line"`
		Points []point `json:"points"`
	}
	sort.Sort(sort.Reverse(sort.IntSlice(frequencies)))
	var series map[point]float64 = make(map[point]float64)
	var points []point
	for i, v := range frequencies {
		points = append(points, point{X: math.Log(float64(i + 1)), Y: math.Log(float64(v))})
		series[points[i]] = float64(v)
	}
	leastSquaresLine := leastSquares(series)
	output := outputStyle{Line: leastSquaresLine, Points: points}
	jsonData, _ := json.Marshal(output)
	ioutil.WriteFile(baseFileName+"\\MessagesByUser_ZipfsPlot.json", jsonData, os.ModePerm)
}

// func numberOfMessagesByLengthCharactersByUsers(messagesByUsers map[string][]userMessage, baseFileName string) {
// 	type outputStyle struct {
// 		User   string            `json:"user"`
// 		Series []lineChartOutput `json:"series"`
// 	}
// 	var numberOfMessagesLengthPerUsers []outputStyle = make([]outputStyle, 0)
// 	maxLength := 0
// 	for _, messages := range messagesByUsers {
// 		messagesContent := extractUserMessageContent(messages)
// 		if findMaxStringLengthCharacters(messagesContent...) > maxLength {
// 			maxLength = findMaxStringLengthCharacters(messagesContent...)
// 		}
// 	}

// 	for user, messages := range messagesByUsers {
// 		var countOfLength []lineChartOutput = make([]lineChartOutput, maxLength+1)
// 		for i := 0; i < len(countOfLength); i++ {
// 			countOfLength[i].X = i
// 		}
// 		for _, message := range messages {
// 			countOfLength[len(message.Content)].Y++
// 		}
// 		numberOfMessagesLengthPerUsers = append(numberOfMessagesLengthPerUsers, outputStyle{User: user, Series: countOfLength})

// 	}
// 	jsonData, _ := json.Marshal(numberOfMessagesLengthPerUsers)
// 	ioutil.WriteFile(baseFileName+"\\NumberOfMessages_LengthCharacters_Users.json", jsonData, os.ModePerm)
// }

// func numberOfMessagesByLengthWordsByUsers(messagesByUsers map[string][]userMessage, baseFileName string) {
// 	type outputStyle struct {
// 		User   string            `json:"user"`
// 		Series []lineChartOutput `json:"series"`
// 	}
// 	var numberOfMessagesLengthPerUsers []outputStyle = make([]outputStyle, 0)
// 	maxLength := 0
// 	for _, messages := range messagesByUsers {
// 		messagesContent := extractUserMessageContent(messages)
// 		if findMaxStringLengthWords(messagesContent...) > maxLength {
// 			maxLength = findMaxStringLengthWords(messagesContent...)
// 		}
// 	}

// 	for user, messages := range messagesByUsers {
// 		var countOfLength []lineChartOutput = make([]lineChartOutput, maxLength+1)
// 		for i := 0; i < len(countOfLength); i++ {
// 			countOfLength[i].X = i
// 		}
// 		for _, message := range messages {
// 			countOfLength[findMaxStringLengthWords(message.Content)].Y++
// 		}
// 		numberOfMessagesLengthPerUsers = append(numberOfMessagesLengthPerUsers, outputStyle{User: user, Series: countOfLength})

// 	}
// 	jsonData, _ := json.Marshal(numberOfMessagesLengthPerUsers)
// 	ioutil.WriteFile(baseFileName+"\\NumberOfMessages_LengthWords_Users.json", jsonData, os.ModePerm)
// }

func allNameChanges(file []string, messageStartRegexp regexp.Regexp, whatsAppNotificationRegexp regexp.Regexp, baseFileName string) {
	groupNames := make([]nameChange, 0)
	for _, line := range file {
		if !messageStartRegexp.MatchString(line) && whatsAppNotificationRegexp.MatchString(line) {
			date := line[0:17]
			time, err := time.Parse("02/01/2006, 15:04", date)
			if err != nil {
				fmt.Printf("There was an error parsing the date %s \n\r", date)
			}
			newName := line[strings.LastIndex(line[:len(line)-1], "\"")+1 : len(line)-1]
			if len(groupNames) > 0 {
				groupNames[len(groupNames)-1].LengthOfTime = math.Round(time.Sub(groupNames[len(groupNames)-1].DateTime).Hours()*100) / 100
			}
			groupNames = append(groupNames, nameChange{DateTime: time, NewName: newName, LengthOfTime: 0})
		}
	}
	groupNames[len(groupNames)-1].LengthOfTime = math.Round(time.Since(groupNames[len(groupNames)-1].DateTime).Hours()*100) / 100
	jsonData, _ := json.Marshal(groupNames)
	ioutil.WriteFile(baseFileName+"\\GroupNames.json", jsonData, os.ModePerm)
}

func dateSummary(messagesByDate map[time.Time][]dateMessage, startDate time.Time, endDate time.Time, baseFileName string) {
	var summaryData []daySummaryInfo
	for rd := rangeDate(startDate, endDate); ; {
		date := rd()
		if date.IsZero() {
			break
		}

		messages, ok := messagesByDate[date]
		if !ok {
			summaryData = append(summaryData, daySummaryInfo{date, 0, 0.0, "", 0})
			continue
		}
		messagesContent := extractDateMessageContent(messages)

		lengthOfLongestMessage := findMaxStringLengthWords(messagesContent...)
		var longestMessageUser string
		for _, message := range messages {
			if len(strings.Split(message.Content, " ")) == lengthOfLongestMessage {
				longestMessageUser = message.User
				break
			}
		}
		summaryData = append(summaryData, daySummaryInfo{date, len(messages), averageNumberOfWords(messagesContent...), longestMessageUser, lengthOfLongestMessage})
	}
	jsonData, _ := json.Marshal(summaryData)
	ioutil.WriteFile(baseFileName+"\\DateSummary.json", jsonData, os.ModePerm)
}

func basicSummary(messagesByDate map[time.Time][]dateMessage, messagesByUser map[string][]userMessage, messages []message, startDate time.Time, endDate time.Time, baseFileName string) {
	messagesContent := extractMessageContent(messages)
	mostPopularMessage := mostCommonString(removeMessages(messagesContent, "<Media omitted>", "This message was deleted")...)
	mostActiveUser := elementWithLargestSlice(messagesByUser)
	result := overallSummaryInfo{MostPopularMessage: mostPopularMessage,
		MostPopularMessageCount:          countOfString(mostPopularMessage, messagesContent),
		NumberOfMessagesSent:             len(messages),
		NumberOfMediaMessages:            countOfString("<Media omitted>", messagesContent),
		NumberOfDeletedMessages:          countOfString("This message was deleted", messagesContent),
		NumberOfDaysWithActivity:         len(messagesByDate),
		NumberOfDaysAnalysed:             int(endDate.Sub(startDate).Hours() / 24),
		MostActiveUser:                   mostActiveUser,
		NumberOfMessagesByMostActiveUser: len(messagesByUser[mostActiveUser])}
	jsonData, _ := json.Marshal(result)
	ioutil.WriteFile(baseFileName+"\\Summary.json", jsonData, os.ModePerm)
}

func numberOfMessagesByTime(messages []message, baseFileName string) {
	var messagesByHour map[int]int = make(map[int]int)
	for _, v := range messages {
		messagesByHour[v.DateTime.Hour()] += 1
	}
	var output []lineChartOutput
	for i := 0; i < 24; i++ {
		num, ok := messagesByHour[i]
		if !ok {
			num = 0
		}
		output = append(output, lineChartOutput{X: i, Y: num})
	}
	jsonData, _ := json.Marshal(output)
	ioutil.WriteFile(baseFileName+"\\NumberOfMessages_Time.json", jsonData, os.ModePerm)
}

func basicSummaryPerUser(messagesByUsers map[string][]userMessage, baseFileName string) {
	var output []userSummary
	for user, messages := range messagesByUsers {
		messagesContent := extractUserMessageContent(messages)
		numberOfMessagesSent := len(messagesContent)
		numberOfMessagesDeleted := countOfString("This message was deleted", messagesContent)
		numberOfMediaMessages := countOfString("<Media omitted>", messagesContent)
		averageLengthOfMessages := averageNumberOfWords(messagesContent...)
		mostCommonMessage := mostCommonString(removeMessages(messagesContent, "<Media omitted>", "This message was deleted")...)
		mostCommonMessageCount := countOfString(mostCommonMessage, messagesContent)
		output = append(output, userSummary{User: user, NumberOfMessagesSent: numberOfMessagesSent, NumberOfMessagesDeleted: numberOfMessagesDeleted, NumberOfMediaMessages: numberOfMediaMessages, AverageMessageLength: averageLengthOfMessages, MostCommonMessage: mostCommonMessage, MostCommonMessageCount: mostCommonMessageCount})
	}
	jsonData, _ := json.Marshal(output)
	ioutil.WriteFile(baseFileName+"\\UserSummary.json", jsonData, os.ModePerm)
}
