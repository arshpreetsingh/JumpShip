package errors

import (
	"errors"
	"fmt"
)

var (
	GreetingText       = "Thanks for trying out"
	ErrInvalidCMD      = "Command '%s' is invalid!. Please type 'help' see available commands"
	ErrInsuffArguments = "Insufficient arguments for '%s'. Expected %d, Got %d"
	ErrInsuffSlot      = "Couldn't create parking lot with %v slots. Please choose 1 or more"
	NoCarFoundByColour = "No vehicles found for colour '%v'"
	NoHistoryFound     = "No %s histories found"
	DuplicateVehicle   = "A car already parked in this registration number '%s'"

	ErrParkingSlotsFull         = errors.New("Sorry, parking lot is full")
	ErrCarNotFound              = errors.New("Not found")
	ErrSlotAlreadyOccupied      = errors.New("Slot: Slot already occupied")
	ErrSlotAlreadyAvailable     = errors.New("Slot: Slot already available")
	ErrInvalidSlotID            = errors.New("Slot: Invalid slot number")
	ErrNoParkingLot             = errors.New("No Parking Lot available, Please create parking lot using 'create_parking_lot <slot-count>'")
	ErrParkingLotAlreadyCreated = errors.New("Parking Lot already created. Try 'status' to get slot availability or 'park_history' to see parking history")
	ErrInvalidTabSpace          = errors.New("Don't use Tab spaces in the commands")
	ErrInvalidInputSlot         = errors.New("Slot: Plese give total slots in numbers")
	ErrInvalidRegNo             = errors.New("Vehicle: Invalid Resgistartion Number")
	ErrEmptyRegNo               = errors.New("Vehicle: Resgistartion number should not be empty")
	ErrEmptyColour              = errors.New("Vehicle: Colour should not be empty")
	ErrInvalidColour            = errors.New("Vehicle: Invalid Colour")
)

// ErrInvalidCommand err wrapper
func ErrInvalidCommand(cmd string) error {
	return fmt.Errorf(ErrInvalidCMD, cmd)
}

// ErrInvalidArguments err wrapper
func ErrInvalidArguments(cmd string, expect, got int) error {
	return fmt.Errorf(ErrInsuffArguments, cmd, expect, got)
}

// ErrInvalidSlotCount err wrapper
func ErrInvalidSlotCount(count int) error {
	return fmt.Errorf(ErrInsuffSlot, count)
}

// ErrNoCarFoundByColour err wrapper
func ErrNoCarFoundByColour(colour string) error {
	return fmt.Errorf(NoCarFoundByColour, colour)
}

// ErrNoHistoyFound err wrapper
func ErrNoHistoyFound(h string) error {
	return fmt.Errorf(NoHistoryFound, h)
}

// ErrDuplicateVehicle err wrapper
func ErrDuplicateVehicle(regNo string) error {
	return fmt.Errorf(DuplicateVehicle, regNo)
}
