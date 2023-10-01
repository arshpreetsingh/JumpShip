package priceplans

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"joi-energy-golang/domain"
	"joi-energy-golang/repository"
)

// value of each price-plan for specific Meter-ID
/*
{"meter-id":{
"price-plan-1":"value/charge",
"price-plan-2":"value/charge",
"price-plan-3":"value/charge",
}}
*/

// Now also work on these tests as well!!
func TestCompareAllPricePlans(t *testing.T) {
	accounts := repository.NewAccounts(map[string]string{"home-sweet-home": "test-plan"})
	// Why passng home-sweet-home here?
	meterReadings := repository.NewMeterReadings(
		map[string][]domain.ElectricityReading{"home-sweet-home": {{
			Time:    time.Now(),
			Reading: 5.0,
		}, {
			Time:    time.Now().Add(-10 * time.Hour),
			Reading: 15.0,
		}}},
	)
	// price-plans
	pricePlans := repository.NewPricePlans(
		[]domain.PricePlan{{
			PlanName: "test-plan",
			UnitRate: 3.0,
		}},
		&meterReadings,
	)
	service := NewService(
		logrus.NewEntry(logrus.StandardLogger()),
		&pricePlans,
		&accounts,
	)
	plans, err := service.CompareAllPricePlans("home-sweet-home")
	expected := domain.PricePlanComparisons{
		PricePlanId: "test-plan",
		PricePlanComparisons: map[string]float64{
			"test-plan": 3.0,
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected.PricePlanId, plans.PricePlanId)
	assert.InDelta(t, expected.PricePlanComparisons["test-plan"], plans.PricePlanComparisons["test-plan"], 0.001)
}

// type MeterReadings struct {
//	meterAssociatedReadings map[string][]domain.ElectricityReading
//}
// []domain.ElectricityReading
// func TestGetWeekData(t *testing.T)  {
// 	//
// 	accounts := repository.NewAccounts(defaultSmartMeterToPricePlanAccounts())
// 	meterReadings := repository.NewMeterReadings(
// 		defaultMeterElectricityReadings(),
// 	)
// 	fmt.Println(accounts, meterReadings)
// 	pricePlans := repository.NewPricePlans(
// 		defaultPricePlans(),
// 		&meterReadings,
// 	)
// 	fmt.Println(pricePlans)
// }

// input: - Smart-meter-id ,
// check for price-plan if exists for specific smart-meter-id?
// pass Plan Values
// create struct //
// use generate_Data() function
// find for some vaulues!!
