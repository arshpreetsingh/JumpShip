package readings

import (
	"testing"

	"github.com/sirupsen/logrus"

	"joi-energy-golang/domain"
	"joi-energy-golang/repository"
//	"fmt"
	//"reflect"
//  "time"
// "reflect"
 "github.com/stretchr/testify/assert"
)

func TestStoreReadings(t *testing.T) {
	meterReadings := repository.NewMeterReadings(
		map[string][]domain.ElectricityReading{},
	)
	service := NewService(
		logrus.NewEntry(logrus.StandardLogger()),
		&meterReadings,
	)
	service.StoreReadings("1", []domain.ElectricityReading{{}})
}

func TestGetReadings(t *testing.T) {
	meterReadings := repository.NewMeterReadings(
		map[string][]domain.ElectricityReading{},
	)
	service := NewService(
		logrus.NewEntry(logrus.StandardLogger()),
		&meterReadings,
	)
	response := service.GetReadings("1")
	var expectedR []domain.ElectricityReading
	//fmt.Println("second",reflect.TypeOf(expectedR))
  assert.Equal(t, expectedR, response)
}
