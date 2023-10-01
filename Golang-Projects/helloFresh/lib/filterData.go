package lib

// Functions to Filter Data based on required values!

import (
	"github.com/arshpreetsingh/hellofresh/src"
	"sort"
	"strings"
)

func uniqueRecipeCount(fileName string, recipeValues chan map[string]int) {
	jsonDecoder := src.ReadJson(fileName)
	filterKeyRecipe := "recipe"
	recipe := src.RecipeCounter(*jsonDecoder, filterKeyRecipe)
	recipeValues <- recipe
}

func busiestPostCode(fileName string, postCode chan string, deliveryCount chan int) {
	//fileName := "hf_test_calculation_fixtures.json"
	//postCodeValues := make(chan map[string]int)
	jsonDecoder := src.ReadJson(fileName)
	filterKeyPostCode := "postcode"
	postCodes := src.PostCodeCounter(*jsonDecoder, filterKeyPostCode)
	keys := make([]string, 0, len(postCodes))

	for key := range postCodes {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return postCodes[keys[i]] > postCodes[keys[j]]
	})

	postCode <- keys[0]
	deliveryCount <- postCodes[keys[0]]
}

func countRelativePostcodeTime(fileName, startTime, endTime, postCode string, deliveryCount chan int) {
	jsonDecoder := src.ReadJson(fileName)
	filterKeyDelivery := "delivery"
	filterKeyPostCode := "postcode"
	deliveryValues := src.DeliveryCounter(*jsonDecoder, filterKeyDelivery, filterKeyPostCode)
	matchCase := startTime + " - " + endTime
	deliveryCount <- deliveryValues[matchCase][postCode]
}

func matchRecipeNames(fileName string, matchList []string, matchedListSorted chan []string) {
	jsonDecoder := src.ReadJson(fileName)
	filterKeyRecipe := "recipe"
	var matchedValues []string
	recipeValues := src.RecipeCounter(*jsonDecoder, filterKeyRecipe)
	for key, _ := range recipeValues {
		ok := func(key string, values []string) bool {
			for _, value := range values {
				//fmt.Println("key", "ValueTOBEMatched")
				//fmt.Println(key, value)
				if strings.Contains(key, value) {
					return true
				}
			}
			return false
		}(key, matchList)
		if ok {
			matchedValues = append(matchedValues, key)
		}
	}
	sort.Strings(matchedValues)
	matchedListSorted <- matchedValues
}
