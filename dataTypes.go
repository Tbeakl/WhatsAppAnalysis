package main

import "time"

type nameChange struct {
	DateTime time.Time `json:"DateTime"`
	NewName  string    `json:"newName"`
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
	NumberOfMesages           int     `json:"NumberOfMessages"`
	AverageMessageLength      float64 `json:"AverageMessageLengthWords"`
	UserWhoSentLongestMessage string  `json:"UserWhoSentLongestMessage"`
	LengthOfLongestMessage    int     `json:"LengthOfLongestMessageWords"`
}
