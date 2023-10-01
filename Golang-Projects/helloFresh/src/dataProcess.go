package src

import (
	"encoding/json"
	"fmt"
	"strings"
)

// RecipeCounter Count the Unique Recipes.
func RecipeCounter(decoder json.Decoder, filterKeyRecipe string) map[string]int {
	data := map[string]interface{}{}
	recipeValues := map[string]int{}
	//logger.Println("Counting Recipe values!")
	for decoder.More() {
		err := decoder.Decode(&data)
		if err != nil {
			//logger.Println("Unable to decode JSON")
		}
		recipeName := fmt.Sprint(data[filterKeyRecipe])
		_, isPresent := recipeValues[recipeName]
		count := 1
		if !isPresent {
			recipeValues[recipeName] = count
		} else {
			//fmt.Println()
			recipeValues[recipeName] = recipeValues[recipeName] + 1
		}

	}
	//logger.Println("Recipe Counter Completed!!")
	return recipeValues
}

//PostCodeCounter Find the postcode with most delivered recipes.
func PostCodeCounter(decoder json.Decoder, filterKeyPostCode string) map[string]int {
	//startTime
	//log.Println("DeliveryCounter: Total Time Taken", time.Since(startTime))
	data := map[string]interface{}{}
	postCodes := map[string]int{}
	//logger.Println("Counting Post-Code values!")
	for decoder.More() {
		err := decoder.Decode(&data)
		if err != nil {
			//logger.Println("Unable to decode JSON")
		}
		postCode := fmt.Sprint(data[filterKeyPostCode])
		_, isPresent := postCodes[postCode]
		count := 1
		if !isPresent {
			postCodes[postCode] = count
		} else {
			//fmt.Println()
			postCodes[postCode] = postCodes[postCode] + 1
		}
	}
	//logger.Println("Post-Codes Counter Completed!")
	return postCodes
}

//4. Count the number of deliveries to postcode
//10120 that lie within the delivery time between 10 am and 3 PM

func DeliveryCounter(decoder json.Decoder, filterKeyDelivery, filterKeyPostCode string) map[string]map[string]int {
	//startTime := time.Now()
	data := map[string]interface{}{}
	deliveryValues := map[string]map[string]int{}
	//logger.Println("DeliveryCounter:Counting Delivery values!")
	for decoder.More() {
		err := decoder.Decode(&data)
		if err != nil {
			//logger.Println("Unable to decode JSON")
		}
		deliveryValueArr := strings.SplitAfter(fmt.Sprint(data[filterKeyDelivery]), " ")
		deliveryValue := strings.Join(deliveryValueArr[1:4], "")
		postCode := fmt.Sprint(data[filterKeyPostCode])
		_, isPresentDelivery := deliveryValues[deliveryValue]
		count := 1
		if !isPresentDelivery {
			postcodeValues2 := map[string]int{}
			postcodeValues2[postCode] = count
			deliveryValues[deliveryValue] = postcodeValues2
		} else {
			deliveryValues[deliveryValue][postCode] = deliveryValues[deliveryValue][postCode] + 1
		}
	}
	//log.Println("DeliveryCounter: Total Time Taken", time.Since(startTime))
	//logger.Println("DeliveryCounter:Delivery counter Completed!")
	return deliveryValues
}
