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

func outputGroupNames(groupNames []nameChange, baseFileName string) {
	// jsonData, _ := json.Marshal(groupNames)
	jsonData, _ := json.MarshalIndent(groupNames, "", "	")
	ioutil.WriteFile(baseFileName+"\\GroupNames.json", jsonData, os.ModePerm)
}

func numberOfMessagesByUsers(messagesByUsers map[string][]message, baseFileName string) {
	var numberOfMessages []barChartOutput
	for user, messages := range messagesByUsers {
		numberOfMessages = append(numberOfMessages, barChartOutput{X: user, Y: len(messages)})
	}

	jsonData, _ := json.Marshal(numberOfMessages)
	ioutil.WriteFile(baseFileName+"\\NumberOfMessages_Users.json", jsonData, os.ModePerm)
}

func numberOfMessagesByLengthCharactersByUsers(messagesByUsers map[string][]message, baseFileName string) {
	type outputStyle struct {
		User   string            `json:"user"`
		Series []lineChartOutput `json:"series"`
	}
	var numberOfMessagesLengthPerUsers []outputStyle = make([]outputStyle, 0)
	maxLength := 0
	for _, messages := range messagesByUsers {
		if findMaxMessageLengthCharacters(messages) > maxLength {
			maxLength = findMaxMessageLengthCharacters(messages)
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
	jsonData, _ := json.Marshal(numberOfMessagesLengthPerUsers)

	ioutil.WriteFile(baseFileName+"\\NumberOfMessages_LengthCharacters_Users.json", jsonData, os.ModePerm)
}

func numberOfMessagesByLengthWordsByUsers(messagesByUsers map[string][]message, baseFileName string) {
	type outputStyle struct {
		User   string            `json:"user"`
		Series []lineChartOutput `json:"series"`
	}
	var numberOfMessagesLengthPerUsers []outputStyle = make([]outputStyle, 0)
	maxLength := 0
	for _, messages := range messagesByUsers {
		if findMaxMessageLengthWords(messages) > maxLength {
			maxLength = findMaxMessageLengthWords(messages)
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
	jsonData, _ := json.Marshal(numberOfMessagesLengthPerUsers)
	ioutil.WriteFile(baseFileName+"\\NumberOfMessages_LengthWords_Users.json", jsonData, os.ModePerm)
}

func allNameChanges(file []string, messageStartRegexp regexp.Regexp, whatsAppNotificationRegexp regexp.Regexp) []nameChange {
	output := make([]nameChange, 0)
	for _, line := range file {
		if !messageStartRegexp.MatchString(line) && whatsAppNotificationRegexp.MatchString(line) {
			date := line[0:17]
			time, err := time.Parse("02/01/2006, 15:04", date)
			if err != nil {
				fmt.Printf("There was an error parsing the date %s \n\r", date)
			}
			newName := line[strings.LastIndex(line[:len(line)-1], "\"")+1 : len(line)-1]
			output = append(output, nameChange{DateTime: time, NewName: newName})
		}
	}
	return output
}

// func daySummary(messagesByUser map[string][]message, startDate time.Time, endDate time.Time) map[time.Time]daySummaryInfo {
// 	output := make(map[time.Time]daySummaryInfo)
// 	for rd := rangeDate(startDate, endDate); ; {
// 		date := rd()

// 	}
// 	return output
// }
