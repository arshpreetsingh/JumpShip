package main

import (
	"fmt"
	"github.com/arshpreetsingh/hellofresh/lib"
	"os"
	"time"
)

//go run main.go hf_test_calculation_fixtures.json 9AM 9PM 10115 pizza banana chicken

func main() {
	logFile := "application.log"
	logger := lib.StartLogs(logFile)
	s := time.Now()
	logger.Println("Application Started", s)
	cmdArgs := os.Args
	fileName := cmdArgs[1]
	startTime := cmdArgs[2]
	endTime := cmdArgs[3]
	postalCode := cmdArgs[4]
	var matchedList []string
	for _, value := range cmdArgs[5:] {
		matchedList = append(matchedList, value)
	}
	user := lib.User{StartTime: startTime, EndTime: endTime, PostalCode: postalCode, MatchedList: matchedList, Filename: fileName}
	finalResult := user.FinalResult()
	fmt.Println(finalResult)
	logger.Println(time.Since(s))
}
