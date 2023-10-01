package readings

import (
	"context"
	"errors"
	"fmt"
	//"time"

	"github.com/go-kit/kit/endpoint"
	validation "github.com/go-ozzo/ozzo-validation"

	"joi-energy-golang/domain"
)

// Should Middleware be present here or any where else?
func makeValidationMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			msg, ok := req.(domain.StoreReadings)
			if !ok {
				return nil, domain.ErrInvalidMessageType
			}
			if err := validateStoreReadings(msg); err != nil {
				return nil, fmt.Errorf("%w: %s", domain.ErrMissingArgument, err)
			}
			return next(ctx, req)
		}
	}
}

func validateStoreReadings(msg domain.StoreReadings) error {
	if err := validation.ValidateStruct(
		&msg,
		validation.Field(&msg.SmartMeterId, validation.Required),
		validation.Field(&msg.ElectricityReadings, validation.NotNil),
	); err != nil {
		return err
	}
	for _, row := range msg.ElectricityReadings {
		if err := validateElectricityReadings(row); err != nil {
			return err
		}
	}
	return nil
}

// This will be updated based on valid Electricit Reading!
// If there is no smart meter-ID it should not show here as well!
func validateElectricityReadings(row domain.ElectricityReading) error {
	if row.Reading <= 0 {
		err := errors.New("Reading should be greater than 0")
		return err
	}

	// if time.Now().Before(row.Time) {
	// 	err := errors.New("TIME should not be greater than Present Time!")
	// 	return err
	// }
	// {2021-06-05 21:05:59 +0530 IST -2}
	return nil
}

func validateSmartMeterId(smartMeterId string) error {
	return validation.Validate(smartMeterId, validation.Required)
}
