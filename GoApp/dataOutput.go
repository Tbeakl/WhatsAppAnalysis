package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

func numberOfMessagesByUsers(messagesByUsers map[string][]userMessage, baseFileName string) {
	var numberOfMessages []barChartOutput
	for user, messages := range messagesByUsers {
		numberOfMessages = append(numberOfMessages, barChartOutput{X: user, Y: len(messages)})
	}

	jsonData, _ := json.MarshalIndent(numberOfMessages, "", "	")
	ioutil.WriteFile(baseFileName+"\\NumberOfMessages_Users.json", jsonData, os.ModePerm)
}

func numberOfMessagesByLengthCharactersByUsers(messagesByUsers map[string][]userMessage, baseFileName string) {
	type outputStyle struct {
		User   string            `json:"user"`
		Series []lineChartOutput `json:"series"`
	}
	var numberOfMessagesLengthPerUsers []outputStyle = make([]outputStyle, 0)
	maxLength := 0
	for _, messages := range messagesByUsers {
		messagesContent := extractUserMessageContent(messages)
		if findMaxStringLengthCharacters(messagesContent...) > maxLength {
			maxLength = findMaxStringLengthCharacters(messagesContent...)
		}
	}

	for user, messages := range messagesByUsers {
		var countOfLength []lineChartOutput = make([]lineChartOutput, maxLength+1)
		for i := 0; i < len(countOfLength); i++ {
			countOfLength[i].X = i
		}
		for _, message := range messages {
			countOfLength[len(message.Content)].Y++
		}
		numberOfMessagesLengthPerUsers = append(numberOfMessagesLengthPerUsers, outputStyle{User: user, Series: countOfLength})

	}
	jsonData, _ := json.MarshalIndent(numberOfMessagesLengthPerUsers, "", "	")
	ioutil.WriteFile(baseFileName+"\\NumberOfMessages_LengthCharacters_Users.json", jsonData, os.ModePerm)
}

func numberOfMessagesByLengthWordsByUsers(messagesByUsers map[string][]userMessage, baseFileName string) {
	type outputStyle struct {
		User   string            `json:"user"`
		Series []lineChartOutput `json:"series"`
	}
	var numberOfMessagesLengthPerUsers []outputStyle = make([]outputStyle, 0)
	maxLength := 0
	for _, messages := range messagesByUsers {
		messagesContent := extractUserMessageContent(messages)
		if findMaxStringLengthWords(messagesContent...) > maxLength {
			maxLength = findMaxStringLengthWords(messagesContent...)
		}
	}

	for user, messages := range messagesByUsers {
		var countOfLength []lineChartOutput = make([]lineChartOutput, maxLength+1)
		for i := 0; i < len(countOfLength); i++ {
			countOfLength[i].X = i
		}
		for _, message := range messages {
			countOfLength[len(strings.Split(message.Content, " "))].Y++
		}
		numberOfMessagesLengthPerUsers = append(numberOfMessagesLengthPerUsers, outputStyle{User: user, Series: countOfLength})

	}
	jsonData, _ := json.MarshalIndent(numberOfMessagesLengthPerUsers, "", "	")
	ioutil.WriteFile(baseFileName+"\\NumberOfMessages_LengthWords_Users.json", jsonData, os.ModePerm)
}

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
			groupNames = append(groupNames, nameChange{DateTime: time, NewName: newName})
		}
	}

	jsonData, _ := json.MarshalIndent(groupNames, "", "	")
	ioutil.WriteFile(baseFileName+"\\GroupNames.json", jsonData, os.ModePerm)
}

func dateSummary(messagesByDate map[time.Time][]dateMessage, startDate time.Time, endDate time.Time, baseFileName string) {
	summaryData := make(map[time.Time]daySummaryInfo)
	for rd := rangeDate(startDate, endDate); ; {
		date := rd()
		if date.IsZero() {
			break
		}

		messages, ok := messagesByDate[date]
		if !ok {
			summaryData[date] = daySummaryInfo{0, 0.0, "", 0}
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
		summaryData[date] = daySummaryInfo{len(messages), averageNumberOfWords(messagesContent...), longestMessageUser, lengthOfLongestMessage}
	}
	jsonData, _ := json.MarshalIndent(summaryData, "", "	")
	ioutil.WriteFile(baseFileName+"\\DateSummary_"+startDate.Format("2006-01-02")+"_"+endDate.Format("2006-01-02")+".json", jsonData, os.ModePerm)
}

func basicSummary(messagesByDate map[time.Time][]dateMessage, messagesByUser map[string][]userMessage, messages []message, startDate time.Time, endDate time.Time, baseFileName string) {
	mostPopularMessage := mostCommonString(extractMessageContent(messages)...)
	mostActiveUser := elementWithLargestSlice(messagesByUser)
	result := overallSummaryInfo{MostPopularMessage: mostPopularMessage,
		MostPopularMessageCount:          countOfString(mostPopularMessage, extractMessageContent(messages)),
		NumberOfMessagesSent:             len(messages),
		NumberOfDaysWithActivity:         len(messagesByDate),
		NumberOfDaysAnalysed:             int(endDate.Sub(startDate).Hours() / 24),
		MostActiveUser:                   mostActiveUser,
		NumberOfMessagesByMostActiveUser: len(messagesByUser[mostActiveUser])}
	jsonData, _ := json.Marshal(result)
	ioutil.WriteFile(baseFileName+"\\Summary.json", jsonData, os.ModePerm)
}
