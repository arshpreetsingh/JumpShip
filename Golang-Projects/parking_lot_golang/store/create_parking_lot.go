package store

import (
	"fmt"
	"strconv"

	"parking_lot/errors"
	"parking_lot/schema"
)

const parkingLotName = "go-jek-lot1"

type createParkingLotStore struct {
	*store
}

// NewCreateParkingLotStore returns new store object
func NewCreateParkingLotStore(st *store) *createParkingLotStore {
	pl := &createParkingLotStore{st}
	return pl
}

func (pl *createParkingLotStore) IsHelp(arg string) (string, bool) {
	if arg == string(schema.CMDHelp) {
		return schema.CMDCreateParkingLotHint, true
	}
	return "", false
}

// Execute - this will create_parking_lot with given slots.
// The system will check if no parking_lot availabe then it create a parking_lot
// with N slots.
// All the slots will initialized with sequence slot numbers by start 1 to N
func (pl *createParkingLotStore) Execute(cmd *schema.Command) (string, error) {
	if res, isHelp := pl.IsHelp(cmd.Arguments[0]); isHelp {
		return res, nil
	}
	totalSlots, err := strconv.Atoi(cmd.Arguments[0])
	if err != nil {
		return "", errors.ErrInvalidInputSlot
	}
	if totalSlots <= 0 {
		return "", errors.ErrInvalidSlotCount(totalSlots)
	}
	if ParkingLot != nil {
		return "", errors.ErrParkingLotAlreadyCreated
	}
	newLot := &schema.ParkingLot{
		Name:        parkingLotName,
		Floor:       "ground_floor",
		TotalBlocks: 1,
		BlockHeight: 12, // feet
		TotalSlots:  totalSlots,
		Slots:       make([]*schema.Slot, totalSlots),
	}

	// initiate nil slot properties
	for i := range newLot.Slots {
		newLot.Slots[i] = new(schema.Slot)
		newLot.Slots[i].SetID(i + 1)
		newLot.Slots[i].SetName(i + 1)
		newLot.Slots[i].BlockID = 1
		newLot.Slots[i].BlockName = "A-Block"
		newLot.Slots[i].MakeSlotFree()
	}

	// set parking lot info global
	ParkingLot = newLot
	return fmt.Sprintf(ParkinglotCreatedInfo, totalSlots), nil
}
