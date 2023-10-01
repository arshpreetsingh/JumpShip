package store

import (
	"fmt"
	"strings"
	"time"

	"parking_lot/errors"
	"parking_lot/schema"
	"parking_lot/utils"
)

const carModelSedan = "sedan"

type parkStore struct {
	*store
}

// NewParkStore returns new store object
func NewParkStore(st *store) *parkStore {
	pl := &parkStore{st}
	return pl
}

func (pl *parkStore) IsHelp(arg string) (string, bool) {
	if arg == string(schema.CMDHelp) {
		return schema.CMDParkHint, true
	}
	return "", false
}

// Execute - `park` Command will takes registration number and colour as Arguments
// the system checks for a first availabe slot to park, if slot available
// slot will allocated to the vehicle.
// This will checks if the vehicle registration number is duplicate or not.
func (pl *parkStore) Execute(cmd *schema.Command) (string, error) {
	if res, isHelp := pl.IsHelp(cmd.Arguments[0]); isHelp {
		return res, nil
	}
	if ParkingLot == nil {
		return "", errors.ErrNoParkingLot
	}
	if err := validateParkReq(cmd.Arguments); err != nil {
		return "", err
	}
	// TODO check for registration number deplication

	// Checks for first available slot
	car := &schema.Vehicle{
		RegistrationNumber: cmd.Arguments[0],
		Colour:             strings.ToLower(cmd.Arguments[1]),
		Type:               string(schema.VehicleTypeCar),
		Model:              carModelSedan,
		Wheels:             4,
		Height:             57, // inches
	}

	availSlot, err := ParkingLot.FirstAvailableSlot()
	if err != nil {
		return "", err
	}
	// park vehicle in the slot
	if err := availSlot.ParkVehicle(car); err != nil {
		return "", err
	}
	parkHistory := &schema.ParkHistory{
		SlotID:             availSlot.GetID(),
		RegistrationNumber: cmd.Arguments[0],
		Colour:             strings.ToLower(cmd.Arguments[1]),
		CreatedAt:          time.Now(),
	}
	// save parking history
	ParkingLot.ParkHistory = append(ParkingLot.ParkHistory, parkHistory)
	err := ParkingLot.CheckDuplicate(car.RegistrationNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(SlotAllocatedInfo, availSlot.GetID()), nil
}

func validateParkReq(args []string) error {
	if !utils.IsRegNoValid(args[0]) {
		return errors.ErrInvalidRegNo
	}
	if !utils.IsValidString(args[1]) {
		return errors.ErrInvalidColour
	}
	return nil
}
