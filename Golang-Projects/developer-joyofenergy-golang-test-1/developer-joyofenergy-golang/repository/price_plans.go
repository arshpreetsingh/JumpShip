package repository

import (
	"time"

	//"reflect"

	"joi-energy-golang/domain"
)

type PricePlans struct {
	pricePlans    []domain.PricePlan
	meterReadings *MeterReadings
}

func NewPricePlans(pricePlans []domain.PricePlan, meterReadings *MeterReadings) PricePlans {
	return PricePlans{
		pricePlans:    pricePlans,
		meterReadings: meterReadings,
	}
}

func filterReadings(readings []domain.ElectricityReading) []domain.ElectricityReading {

	timeObjectWeekly := time.Now().AddDate(0, 0, -7)
	var weeklyreadings []domain.ElectricityReading
	for _, value := range readings {
		if timeObjectWeekly.After(value.Time) {
			weeklyreadings = append(weeklyreadings, value)
		}
	}
	return weeklyreadings
}

func (p *PricePlans) ConsumptionElectricityCostWeekly(smartMeterId string) map[string]float64 {
	electricityReadings := p.meterReadings.GetReadings(smartMeterId)
	electricityReadingsWeekly := filterReadings(electricityReadings)
	if len(electricityReadingsWeekly) == 0 {
		return map[string]float64{}
	}
	costs := map[string]float64{}
	for _, plan := range p.pricePlans {
		costs[plan.PlanName] = calculateCost(electricityReadingsWeekly, plan)
	}
	return costs
}

func (p *PricePlans) ConsumptionCostOfElectricityReadingsForEachPricePlan(smartMeterId string) map[string]float64 {
	electricityReadings := p.meterReadings.GetReadings(smartMeterId)
	costs := map[string]float64{}
	for _, plan := range p.pricePlans {
		costs[plan.PlanName] = calculateCost(electricityReadings, plan)
	}
	return costs
}

func calculateCost(electricityReadings []domain.ElectricityReading, pricePlan domain.PricePlan) float64 {
	// average cost of specific Price_Plan
	average := calculateAverageReading(electricityReadings)
	timeElapsed := calculateTimeElapsed(electricityReadings)
	averagedCost := average / timeElapsed.Hours()
	return averagedCost * pricePlan.UnitRate
}

func calculateAverageReading(electricityReadings []domain.ElectricityReading) float64 {
	sum := 0.0
	for _, r := range electricityReadings {
		sum += r.Reading
	}
	return sum / float64(len(electricityReadings))
}

// Need to UnderStand This!! For Sure!!!
// Need to understnad this!! For Sure at all levels!!! yes Yes! yes!!!!
func calculateTimeElapsed(electricityReadings []domain.ElectricityReading) time.Duration {
	// Meaning is like How much time has been passed between First and last!!
	// Time Range, in mintues.
	var first, last time.Time
	for _, r := range electricityReadings {
		if r.Time.Before(first) || (first == time.Time{}) {
			first = r.Time
			//fmt.Println("this is first", first)
		}
	}
	for _, r := range electricityReadings {
		if r.Time.After(last) || (last == time.Time{}) {
			last = r.Time
			//fmt.Println("this is last", last)
		}
	}
	return last.Sub(first)
}
