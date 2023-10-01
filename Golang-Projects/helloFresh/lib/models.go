package lib

type finalResult struct {
	UniqueRecipeCount int            `json:"unique_recipe_count"`
	CountPerRecipe    map[string]int `json:"count_per_recipe"`
	BusiestPostcode   struct {
		Postcode      string `json:"postcode"`
		DeliveryCount int    `json:"delivery_count"`
	} `json:"busiest_postcode"`
	CountPerPostcodeAndTime struct {
		Postcode      string `json:"postcode"`
		From          string `json:"from"`
		To            string `json:"to"`
		DeliveryCount int    `json:"delivery_count"`
	} `json:"count_per_postcode_and_time"`
	MatchByName []string `json:"match_by_name"`
}
