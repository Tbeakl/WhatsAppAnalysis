package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	filename := os.Args[1]
	textlines := readInFile(filename + ".txt")

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

	var groupName []nameChange = allNameChanges(textlines, *messageStartRegexp, *whatsAppNotificationRegexp)

	//Some further data cleaning and processing
	textlines = removeAllNotifications(textlines, *messageStartRegexp, *whatsAppNotificationRegexp)

	messagesByUser := splitByUsers(textlines, *messageStartRegexp, *whatsAppNotificationRegexp)

	outputGroupNames(groupName, filename)
	numberOfMessagesByUsers(messagesByUser, filename)
	numberOfMessagesByLengthCharactersByUsers(messagesByUser, filename)
	numberOfMessagesByLengthWordsByUsers(messagesByUser, filename)
}
