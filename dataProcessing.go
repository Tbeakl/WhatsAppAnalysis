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

func makeUsernameConsistent(file []string, messageStartRegexp regexp.Regexp, whatsAppNotificationRegexp regexp.Regexp) []string {
	usernameChanger := make(map[string]string)
	var output []string = make([]string, 0)
	for i := len(file) - 1; i >= 0; i-- {
		line := file[i]
		if !messageStartRegexp.MatchString(line) && whatsAppNotificationRegexp.MatchString(line) {
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
		} else if messageStartRegexp.MatchString(line) {
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

func splitByUsers(file []string, messageStartRegexp regexp.Regexp, whatsAppNotificationRegexp regexp.Regexp) map[string][]message {
	lastLine := ""
	var messagesByUser map[string][]message = make(map[string][]message)

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
			messageContent := line[indexOfSetInput+1:] + lastLine
			messagesByUser[user] = append(messagesByUser[user], message{time, messageContent})

			lastLine = ""
		} else {
			//This means it can go on the last line as a prepend
			lastLine = line + lastLine
		}
	}

	return messagesByUser
}
