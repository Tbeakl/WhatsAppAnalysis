package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

//This reads in a given textfile, outputting each line into a separate element of the string slice
func readInFile(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textlines []string

	for scanner.Scan() {
		textlines = append(textlines, scanner.Text())
	}

	file.Close()
	return textlines
}

//This makes sure that each person is only represented once even if they change phonenumber, provided WhatsApp gave the notification that their number changed.
func makeUsernameConsistent(file []string, whatsAppNotificationRegexp regexp.Regexp) []string {
	usernameChanger := make(map[string]string)
	var output []string = make([]string, 0)
	for i := len(file) - 1; i >= 0; i-- {
		line := file[i]
		isNotification := whatsAppNotificationRegexp.MatchString(line)
		if isNotification && strings.LastIndex(line, ":") == 14 {
			//This means that we have a notification from WhatsApp
			if strings.Contains(line, " changed to ") {
				oldName := strings.Split(line, " changed to ")[0][20:]
				newName := strings.Split(line, " changed to ")[1]
				oldName = strings.ReplaceAll(oldName, " ", " ") //These are changing nonbreaking spaces in to normal spaces
				newName = strings.ReplaceAll(newName, " ", " ")
				_, changed := usernameChanger[newName]
				if changed {
					newName = usernameChanger[newName]
				}
				usernameChanger[oldName] = newName
			}
		} else if isNotification {
			indexOfSetInput := 20 + strings.Index(line[20:], ":")
			user := line[20:indexOfSetInput]
			newUser, ok := usernameChanger[user]
			if ok {
				line = strings.Replace(line, user, newUser, 1)
			}
		}
		output = append([]string{line}, output...)
	}
	return output
}

//This removes all notifications from the chat except the ones saying the group name changed
func removeUnusedNotifications(file []string, messageStartRegexp regexp.Regexp, whatsAppNotificationRegexp regexp.Regexp) []string {
	var output []string = make([]string, 0)
	for _, line := range file {
		if !messageStartRegexp.MatchString(line) && whatsAppNotificationRegexp.MatchString(line) {
			if strings.Contains(line, " changed the subject from ") || strings.Contains(line, " created group ") {
				output = append(output, line)
			}
		} else {
			output = append(output, line)
		}
	}
	return output
}

//This removes all notifications from the chat
func removeAllNotifications(file []string, messageStartRegexp regexp.Regexp, whatsAppNotificationRegexp regexp.Regexp) []string {
	var output []string = make([]string, 0)
	for _, line := range file {
		if messageStartRegexp.MatchString(line) || !whatsAppNotificationRegexp.MatchString(line) {
			output = append(output, line)
		}
	}
	return output
}

//This changes the filtered list of messages into an actual slice of messages
func makeIntoMessages(file []string, messageStartRegexp regexp.Regexp) []message {
	lastLine := ""
	var messages []message = make([]message, 0)

	for i := len(file) - 1; i >= 0; i-- {
		line := file[i]

		if messageStartRegexp.MatchString(line) {
			date := line[0:17]
			time, err := time.Parse("02/01/2006, 15:04", date)
			if err != nil {
				fmt.Printf("There was an error parsing the date %s \n\r", date)
			}

			indexOfSetInput := 20 + strings.Index(line[20:], ":")
			user := line[20:indexOfSetInput]
			messageContent := line[indexOfSetInput+2:] + lastLine
			messages = append(messages, message{time, messageContent, user})

			lastLine = ""
		} else {
			//This means it can go on the last line as a prepend
			lastLine = line + lastLine
		}
	}
	return messages
}

//This turns a message slice into a map where you can efficiently lookup by user
func splitByUsers(messages []message) map[string][]userMessage {
	var messagesByUser map[string][]userMessage = make(map[string][]userMessage)

	for _, message := range messages {
		messagesByUser[message.User] = append(messagesByUser[message.User], userMessage{message.DateTime, message.Content})
	}

	return messagesByUser
}

//This turns a message slice into a map where you can efficiently lookup by date
func splitByDay(messages []message) map[time.Time][]dateMessage {
	output := make(map[time.Time][]dateMessage)
	for _, message := range messages {
		year, month, day := message.DateTime.Date()
		date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		output[date] = append(output[date], dateMessage{message.Content, message.User})
	}
	return output
}
