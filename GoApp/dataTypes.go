package main

import "time"

type nameChange struct {
	DateTime     time.Time `json:"Date"`
	NewName      string    `json:"Name"`
	LengthOfTime float64   `json:"LengthOfTime"`
}

type userMessage struct {
	DateTime time.Time `json:"DateTime"`
	Content  string    `json:"Content"`
}

type message struct {
	DateTime time.Time `json:"DateTime"`
	Content  string    `json:"Content"`
	User     string    `json:"User"`
}

type dateMessage struct {
	Content string `json:"Content"`
	User    string `json:"User"`
}

type barChartOutput struct {
	X string `json:"x"`
	Y int    `json:"y"`
}

type lineChartOutput struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type daySummaryInfo struct {
	Date                      time.Time `json:"Date"`
	NumberOfMesages           int       `json:"NumberOfMessages"`
	AverageMessageLength      float64   `json:"AverageMessageLengthWords"`
	UserWhoSentLongestMessage string    `json:"UserWhoSentLongestMessage"`
	LengthOfLongestMessage    int       `json:"LengthOfLongestMessageWords"`
}

type overallSummaryInfo struct {
	MostPopularMessage               string  `json:"MostPopularMessage"`
	MostPopularMessageCount          int     `json:"MostPopularMessageCount"`
	NumberOfMediaMessages            int     `json:"NumberOfMediaMessages"`
	NumberOfDeletedMessages          int     `json:"NumberOfDeletedMessages"`
	NumberOfMessagesSent             int     `json:"NumberOfMessagesSent"`
	AverageMessageLength             float64 `json:"AverageMessageLength"`
	NumberOfDaysWithActivity         int     `json:"NumberOfDaysWithActivity"`
	NumberOfDaysAnalysed             int     `json:"NumberOfDaysAnalysed"`
	MostActiveUser                   string  `json:"MostActiveUser"`
	NumberOfMessagesByMostActiveUser int     `json:"NumberOfMessagesByMostActiveUser"`
}

type line struct {
	Gradient   float64 `json:"m"`
	YIntercept float64 `json:"c"`
}

type point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type userSummary struct {
	User                    string  `json:"User"`
	NumberOfMessagesSent    int     `json:"NumberOfMessagesSent"`
	NumberOfMessagesDeleted int     `json:"NumberOfMessagesDeleted"`
	NumberOfMediaMessages   int     `json:"NumberOfMediaMessages"`
	AverageMessageLength    float64 `json:"AverageMessageLength"`
	MostCommonMessage       string  `json:"MostCommonMessage"`
	MostCommonMessageCount  int     `json:"MostCommonMessageCount"`
}
