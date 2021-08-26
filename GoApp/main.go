package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

func main() {
	fmt.Println("Started")
	filename := os.Args[1]
	textlines := readInFile(filename + ".txt")
	fmt.Println("Read In File")
	filename = "Data_" + filename
	os.Mkdir(filename, os.ModePerm)
	fmt.Println("Created Directory")
	messageStartRegexp, err := regexp.Compile(`\d\d/\d\d/\d\d\d\d, \d\d:\d\d - .*:.*`)
	if err != nil {
		fmt.Printf("An error was returned in the message start compilation %s", err)
	}

	whatsAppNotificationRegexp, err := regexp.Compile(`\d\d/\d\d/\d\d\d\d, \d\d:\d\d -.*`)
	if err != nil {
		fmt.Printf("An error was returned in the message notification compilation %s", err)
	}
	fmt.Println("Compiled Regexp")
	//Here is the initial data preprocessing, this is where the application is the slowest
	textlines = makeUsernameConsistent(textlines, *messageStartRegexp, *whatsAppNotificationRegexp)
	fmt.Println("Made names consistent")
	textlines = removeUnusedNotifications(textlines, *messageStartRegexp, *whatsAppNotificationRegexp)
	fmt.Println("Removed the unused whatsapp notifications")
	allNameChanges(textlines, *messageStartRegexp, *whatsAppNotificationRegexp, filename)
	fmt.Println("Found all group names")
	//Some further data cleaning and processing
	textlines = removeAllNotifications(textlines, *messageStartRegexp, *whatsAppNotificationRegexp)
	fmt.Println("Removed all remaining notifications")
	messages := makeIntoMessages(textlines, *messageStartRegexp)
	fmt.Println("Converted into messages")
	messagesByUser := splitByUsers(messages)
	fmt.Println("Converted into messages by user")
	messagesByDate := splitByDay(messages)
	fmt.Println("Converted into messages by date")
	numberOfMessagesByTime(messages, filename)
	fmt.Println("Calculated number of messages per hour in the day")
	dateSummary(messagesByDate, time.Date(2020, 10, 01, 0, 0, 0, 0, time.UTC), time.Now(), filename)
	fmt.Println("Made the date summary")
	basicSummary(messagesByDate, messagesByUser, messages, messages[0].DateTime, messages[len(messages)-1].DateTime, filename)
	fmt.Println("Made the basic summary")
	basicSummaryPerUser(messagesByUser, filename)
	fmt.Println("Made the user summaries")
}
