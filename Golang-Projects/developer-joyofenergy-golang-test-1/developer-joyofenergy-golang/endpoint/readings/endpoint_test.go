package readings

import (
	"context"
	"testing"
  "fmt"
//	 "reflect"

	"github.com/stretchr/testify/assert"

	"joi-energy-golang/domain"
)


func TestStoreReadingsReturnResultFromService(t *testing.T) {
	s := &MockService{}
	e := makeStoreReadingsEndpoint(s)

  // Need to pass some vlaues to this store-Readings!
	response, err := e(context.Background(), domain.StoreReadings{})
	expectedResponse := domain.StoreReadings{}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)

}

func TestGetReadingEndpointService(t *testing.T) {
	s := &MockService{}
	e:= makeGetReadingsEndpoint(s)

	response, err := e(context.Background(), "123")
	//var expectedR []domain.ElectricityReading
	expectedResponse:=domain.StoreReadings{}
	expectedResponse.SmartMeterId = "123"

	fmt.Println("response",response,err)
	fmt.Println("expected-response",expectedResponse)

	// assert.NoError(t, err)
	// assert.Equal(t, expectedResponse, response)

}
