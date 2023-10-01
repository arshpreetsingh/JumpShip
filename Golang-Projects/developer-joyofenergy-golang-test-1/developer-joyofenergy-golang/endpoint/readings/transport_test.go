package readings

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
//	"reflect"
 "fmt"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/mock"

	"joi-energy-golang/domain"
)

type MockService struct{
	mock.Mock
}


func (s *MockService) StoreReadings(smartMeterId string, reading []domain.ElectricityReading) {}

func (s *MockService) GetReadings(smartMeterId string) []domain.ElectricityReading {
	return []domain.ElectricityReading{}
}


func TestMakeStoreReadingsHandler(t *testing.T) {
	//mockService := new(MockService)
	mockService := &MockService{}
	mockLogger := logrus.New().WithField("test", "mock")
	//fmt.Println("heyyyyy Broooooo",reflect.TypeOf(mockService.Service))
	h := MakeStoreReadingsHandler(mockService, mockLogger)
	r := httptest.NewRecorder()

	input := generateValidInput()
	buf := bytes.NewBuffer(nil)
	data, _ := json.MarshalIndent(&input, "", "  ")
	buf.Write(data)

	req := httptest.NewRequest("POST", "/Endpoint", buf)
	req.Header.Set("Content-type", "application/json")

	h.ServeHTTP(r, req)

	result := r.Result()
	actualStatusCode := result.StatusCode
	assert.Equal(t, http.StatusOK, actualStatusCode)
	err := result.Body.Close()
	assert.NoError(t, err)
}

func TestMakeStoreReadingsHandlerWithInvalidInput(t *testing.T) {
	mockService := &MockService{}
	mockLogger := logrus.New().WithField("test", "mock")
	h := MakeStoreReadingsHandler(mockService, mockLogger)
	r := httptest.NewRecorder()

	req := httptest.NewRequest("POST", "/Endpoint", nil)
	req.Header.Set("Content-type", "application/json")

	h.ServeHTTP(r, req)

	result := r.Result()
	actualStatusCode := result.StatusCode
	assert.Equal(t, http.StatusInternalServerError, actualStatusCode)

	expectedMessage := domain.Error{
		ErrorMessage: "unexpected end of JSON input",
	}
	expected, _ := json.MarshalIndent(expectedMessage, "", "  ")
	actual, err := ioutil.ReadAll(result.Body)
	_ = result.Body.Close()

	assert.NoError(t, err)
	assert.Equal(t, string(expected), string(actual))
}

func TestMakeGetReadingsHandler(t *testing.T) {
	mockService := &MockService{}
	mockLogger := logrus.New().WithField("test", "mock")
	h := MakeGetReadingsHandler(mockService, mockLogger)
	r := httptest.NewRecorder()
	// need to check how to pass valid input to GET Function.
  req := httptest.NewRequest("GET", "/readings/read/smart-meter-1234", nil)
	req.Header.Set("Content-type", "application/json")

	h.ServeHTTP(r, req)

	result := r.Result()
	//fmt.Println("this is result!",result)
	actualStatusCode := result.StatusCode
	assert.Equal(t, http.StatusOK, actualStatusCode)
	err := result.Body.Close()
	assert.NoError(t, err)
}

func TestMakeGetReadingsHandlerInvalid(t *testing.T) {
	mockService := &MockService{}
	mockLogger := logrus.New().WithField("test", "mock")
	h := MakeGetReadingsHandler(mockService, mockLogger)
	r := httptest.NewRecorder()
	// need to check how to pass valid input to GET Function.
  req := httptest.NewRequest("GET", "/readings/read/", nil)
	req.Header.Set("Content-type", "application/json")

	h.ServeHTTP(r, req)

	result := r.Result()
	fmt.Println("Result!!",result)
	//fmt.Println("this is result!",result)
	actualStatusCode := result.StatusCode
	assert.Equal(t, http.StatusInternalServerError, actualStatusCode)
	err := result.Body.Close()
	expectedMessage := domain.Error{
		ErrorMessage: "cannot be blank",
	}
	expected, _ := json.MarshalIndent(expectedMessage, "", "  ")
	actual, err := ioutil.ReadAll(result.Body)
	_ = result.Body.Close()

	assert.NoError(t, err)
	assert.Equal(t, string(expected), string(actual))
}
