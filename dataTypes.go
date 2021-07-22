package main

import "time"

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

type daySummaryInfo struct {
	numberOfMesages           int     `json:"numberOfMessages"`
	averageMessageLength      float64 `json:"averageMessageLength"`
	userWhoSentMostMessages   string  `json:"userWhoSentMostMessages"`
	userWhoSentLongestMessage string  `json:"userWhoSentLongestMessage"`
}
