package schema

import (
	"parking_lot/errors"
	"time"
)

// ParkingLot struct holds all the parking lot information and parking history
type ParkingLot struct {
	Name        string         `json:"name"`
	Floor       string         `json:"floor"`
	TotalBlocks int            `json:"total_blocks"`
	BlockHeight int            `json:"block_height"`
	TotalSlots  int            `json:"total_slots"`
	Address     string         `json:"address"`
	Pincode     string         `json:"pincode"`
	Slots       []*Slot        `json:"slots"`
	ParkHistory []*ParkHistory `json:"park_history"`
}

// ParkHistory holds the parking information
type ParkHistory struct {
	SlotID             uint
	RegistrationNumber string
	Colour             string
	CreatedAt          time.Time
}

// FirstAvailableSlot returns the first available slot to park Vehicle
func (pl *ParkingLot) FirstAvailableSlot() (*Slot, error) {
	for _, slot := range pl.Slots {
		if slot.IsSlotAvailable() {
			return slot, nil
		}
	}

	return nil, errors.ErrParkingSlotsFull
}
func (pl *ParkingLot) CheckDuplicate(RegistratinNo string) error {
	for _, slot := range pl.Slots {
		if RegistratinNo == slot.Vehicle.RegistrationNumber {
			return errors.ErrDuplicateVehicle(RegistratinNo)
		}
	}
	return nil
}
