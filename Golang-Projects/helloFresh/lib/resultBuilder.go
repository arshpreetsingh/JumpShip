package lib

import (
	"encoding/json"
	"fmt"
)

type User struct {
	StartTime   string
	EndTime     string
	PostalCode  string
	MatchedList []string
	Filename    string
}

func (user *User) FinalResult() string {
	//1
	recipeValues := make(chan map[string]int)
	go uniqueRecipeCount(user.Filename, recipeValues)

	//2
	postCode := make(chan string)
	deliveryCount := make(chan int)
	go busiestPostCode(user.Filename, postCode, deliveryCount)

	//3
	deliveryValueCount := make(chan int)
	go countRelativePostcodeTime(user.Filename, user.StartTime, user.EndTime, user.PostalCode, deliveryValueCount)

	//4
	matchList := []string{"Chicken", "Chops", "Pizzas"}
	matchedListSorted := make(chan []string)
	go matchRecipeNames(user.Filename, matchList, matchedListSorted)
	//var final src.FinalResult
	//&final.UniqueRecipeCount := len(<-recpieValues)
	//fmt.Println(final.UniqueRecipeCount)
	var f finalResult
	recipeCounts := <-recipeValues
	f.UniqueRecipeCount = len(recipeCounts)
	f.CountPerRecipe = recipeCounts
	f.BusiestPostcode.Postcode = <-postCode
	f.BusiestPostcode.DeliveryCount = <-deliveryCount
	f.CountPerPostcodeAndTime.Postcode = user.PostalCode
	f.CountPerPostcodeAndTime.From = user.StartTime
	f.CountPerPostcodeAndTime.To = user.EndTime
	f.CountPerPostcodeAndTime.DeliveryCount = <-deliveryValueCount
	f.MatchByName = <-matchedListSorted
	result := createJson(f)
	return result
}

// Create JSON from given Struct!
func createJson(result finalResult) string {
	cASByte, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Not able to Create JSON from Result, Invalid Data")
	}
	return string(cASByte)
}
