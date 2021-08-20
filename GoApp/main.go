package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

func main() {
	filename := os.Args[1]
	textlines := readInFile(filename + ".txt")
	filename = "Data_" + filename
	os.Mkdir(filename, os.ModePerm)
	messageStartRegexp, err := regexp.Compile(`\d\d/\d\d/\d\d\d\d, \d\d:\d\d - .*:.*`)
	if err != nil {
		fmt.Printf("An error was returned in the message start compilation %s", err)
	}

	whatsAppNotificationRegexp, err := regexp.Compile(`\d\d/\d\d/\d\d\d\d, \d\d:\d\d -.*`)
	if err != nil {
		fmt.Printf("An error was returned in the message notification compilation %s", err)
	}

	//Here is the initial data preprocessing
	textlines = makeUsernameConsistent(textlines, *messageStartRegexp, *whatsAppNotificationRegexp)
	textlines = removeUnusedNotifications(textlines, *messageStartRegexp, *whatsAppNotificationRegexp)

	allNameChanges(textlines, *messageStartRegexp, *whatsAppNotificationRegexp, filename)

	//Some further data cleaning and processing
	textlines = removeAllNotifications(textlines, *messageStartRegexp, *whatsAppNotificationRegexp)
	messages := makeIntoMessages(textlines, *messageStartRegexp)
	messagesByUser := splitByUsers(messages)
	messagesByDate := splitByDay(messages)
	numberOfMessagesByTime(messages, filename)
	dateSummary(messagesByDate, time.Date(2020, 10, 01, 0, 0, 0, 0, time.UTC), time.Now(), filename)
	basicSummary(messagesByDate, messagesByUser, messages, messages[0].DateTime, messages[len(messages)-1].DateTime, filename)
	basicSummaryPerUser(messagesByUser, filename)
}
