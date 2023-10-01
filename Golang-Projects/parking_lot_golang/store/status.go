package store

import (
	"fmt"
	"strings"

	"parking_lot/errors"
	"parking_lot/schema"
	"parking_lot/utils"
)

type statusStore struct {
	*store
}

// NewStatusStore returns new store object
func NewStatusStore(st *store) *statusStore {
	pl := &statusStore{st}
	return pl
}

// Execute will returns the current status of all the slots.
func (pl *statusStore) Execute(cmd *schema.Command) (string, error) {
	if ParkingLot == nil {
		return "", errors.ErrNoParkingLot
	}
	var slotStatus = []string{fmt.Sprintf("%-10s%-20s%-10s", "Slot No.", "Registration No", "Colour")}
	for _, slot := range ParkingLot.Slots {
		if slot.IsFree {
			slotStatus = append(slotStatus, fmt.Sprintf("%-10d%-20s%-10s", slot.GetID(), "Slot is free", ""))
		} else {
			slotStatus = append(slotStatus, fmt.Sprintf("%-10d%-20s%-10s", slot.GetID(),
				slot.Vehicle.GetRegNumber(), slot.Vehicle.GetColour()))
		}
	}
	return strings.Join(slotStatus, utils.NewLineDelim), nil
}
