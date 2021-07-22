package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	filename := os.Args[1]
	textlines := readInFile(filename + ".txt")

	var messagesByUser map[string]([]message) = make(map[string]([]message))
	var groupName []nameChange = make([]nameChange, 0)
	var userNameChanger map[string]string = make(map[string]string)

	lastLine := ""

	for i := len(textlines) - 1; i >= 0; i-- {
		line := textlines[i]
		messageStart, err := regexp.MatchString("\\d\\d/\\d\\d/\\d\\d\\d\\d, \\d\\d:\\d\\d - .*:.*", line)
		if err != nil {
			fmt.Printf("Error on regex with string %s \n\r", line)
			continue
		}

		if messageStart {
			date := line[0:17]
			time, err := time.Parse("02/01/2006, 15:04", date)
			if err != nil {
				fmt.Printf("There was an error parsing the date %s \n\r", date)
			}
			indexOfSetInput := 20 + strings.Index(line[20:], ":")
			user := line[20:indexOfSetInput]

			newUser, ok := userNameChanger[user]
			if ok {
				user = newUser
			}

			messageContent := line[indexOfSetInput+1:] + lastLine

			// var curMessage message = message{time, user, messageContent}
			messagesByUser[user] = append(messagesByUser[user], message{time, messageContent})

			lastLine = ""
		} else {
			whatsAppNotification, err := regexp.MatchString("\\d\\d/\\d\\d/\\d\\d\\d\\d, \\d\\d:\\d\\d -.*", line)
			if err != nil {
				fmt.Printf("Error on regex with string %s \n\r", line)
				continue
			}

			if whatsAppNotification {
				switch {
				case strings.Contains(line, " changed to "):
					oldName := strings.Split(line, " changed to ")[0][20:]
					newName := strings.Split(line, " changed to ")[1]
					oldName = strings.ReplaceAll(oldName, " ", " ") //These are changing nonbreaking spaces in to normal spaces
					newName = strings.ReplaceAll(newName, " ", " ")
					_, changed := userNameChanger[newName]
					if changed {
						newName = userNameChanger[newName]
					}
					userNameChanger[oldName] = newName
				case strings.Contains(line, " changed the subject from "), strings.Contains(line, " created group "):
					date := line[0:17]
					time, err := time.Parse("02/01/2006, 15:04", date)
					if err != nil {
						fmt.Printf("There was an error parsing the date %s \n\r", date)
					}
					newName := line[strings.LastIndex(line[:len(line)-1], "\"")+1 : len(line)-1]
					groupName = append(groupName, nameChange{DateTime: time, NewName: newName})
				default:
				}
			} else {
				//This means that it can go into last line as a prepend
				lastLine = line + lastLine
			}
		}
	}

	lastLine = ""

	outputGroupNames(groupName, filename)
	numberOfMessagesByUsers(messagesByUser, filename)
	numberOfMessagesByLengthCharactersByUsers(messagesByUser, filename)
	numberOfMessagesByLengthWordsByUsers(messagesByUser, filename)
}

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

func findMaxMessageLengthCharacters(messages []message) int {
	maxLength := 0
	for _, element := range messages {
		if len(element.Content) > maxLength {
			maxLength = len(element.Content)
		}
	}
	return maxLength
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

func findMaxMessageLengthWords(messages []message) int {
	maxLength := 0
	for _, element := range messages {
		if len(strings.Split(element.Content, " ")) > maxLength {
			maxLength = len(strings.Split(element.Content, " "))
		}
	}
	return maxLength
}

type nameChange struct {
	DateTime time.Time `json:"DateTime"`
	NewName  string    `json:"newName"`
}

type message struct {
	DateTime time.Time `json:"DateTime"`
	Content  string    `json:"Content"`
}

type barChartOutput struct {
	X string `json:"x"`
	Y int    `json:"y"`
}

type lineChartOutput struct {
	X int `json:"x"`
	Y int `json:"y"`
}
