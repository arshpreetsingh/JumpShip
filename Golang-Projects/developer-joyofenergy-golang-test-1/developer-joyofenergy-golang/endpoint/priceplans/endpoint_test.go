package priceplans

import (
	"context"
	"errors"
	"testing"
	//"fmt"

	"github.com/stretchr/testify/assert"

	"joi-energy-golang/domain"
)

func TestCompareAllPricePlansReturnResultFromService(t *testing.T) {
	s := &MockService{}
	e := makeCompareAllPricePlansEndpoint(s)

  // here this will save us, Need to understnad
	response, err := e(context.Background(), "123")
	expectedResponse := domain.PricePlanComparisons{}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}
//
// func TestRecommendPricePlansReturnResultFromService(t *testing.T) {
// 	s:=&MockService{}
// 	e:=makeRecommendPricePlansEndpoint(s)
// 	//makeRecommendPricePlansEndpoint(s Service) endpoint.Endpoint {
// 	// var requestData
// 	// requestData
//
// 	// type PricePlans struct {
// 	// 	pricePlans    []domain.PricePlan
// 	// 	meterReadings *MeterReadings
// 	// }
// 	// here this will save us, Need to understand
// 	// this is video need to be watch before moving further:
// 	// https://www.youtube.com/watch?v=h2RdcrMLQAo
// 	ctx:= context.Background()
// 	ctx = context.WithValue(ctx, "url.Values", "10")
// 	//limitString := ctx.Value(contextkeys.QueryValues).(url.Values).Get("limit")
// 	response, err := e(ctx, domain.PricePlan{})
// 	expectedResponse := domain.PricePlan{}
// 	// expected: domain.StoreReadings{SmartMeterId:"", ElectricityReadings:[]domain.ElectricityReading(nil)}
// 	// actual  : domain.StoreReadings{SmartMeterId:"123", ElectricityReadings:[]domain.ElectricityReading{}}
//   fmt.Println("this is response",response)
// 	fmt.Println("this is expected-response",expectedResponse)
// 	fmt.Println("this is error",err)
//
// 	//assert.NoError(t, err)
// 	//assert.Equal(t, expectedResponse, response)
//
// 	}

func TestCompareAllPricePlansHandleServiceError(t *testing.T) {
	s := &MockService{err: errors.New("oops")}
	e := makeCompareAllPricePlansEndpoint(s)

	_, err := e(context.Background(), "123")
	expectedErr := "oops"

	assert.EqualError(t, err, expectedErr)
}
