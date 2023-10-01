package domain

import (
	"time"
)

type ElectricityReading struct {
	Time    time.Time
	Reading float64
}
//{Dr Evil's Dark Energy price-plan-0 10 []}
type PricePlan struct {
	EnergySupplier      string
	PlanName            string
	UnitRate            float64
	PeakTimeMultipliers []PeakTimeMultiplier
}

type PeakTimeMultiplier struct {
	DayOfWeek  time.Weekday
	Multiplier float64
}

// It should be {Price_plan,Average_Cost_per_hour}
//
type SingleRecommendation struct {
	Key   string
	Value float64
}

type PricePlanRecommendation struct {
	Recommendations []SingleRecommendation
}

type PricePlanComparisons struct {
	PricePlanId          string
	PricePlanComparisons map[string]float64
}

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Error struct {
	ErrorMessage string `json:"errorMessage"`
}

type Message struct {
	ID   string `json:"id"`
	Data string `json:"data"`
	Rows []string `json:"rows"`
}

type Response struct {
}

type StoreReadings struct {
	SmartMeterId        string `json:"smartMeterId"`
	ElectricityReadings []ElectricityReading `json:"electricityReadings"`
}
