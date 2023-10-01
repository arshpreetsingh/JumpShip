package store

import (
	"parking_lot/schema"
)

// Store interface holds all the available cmd exc methods
type Store interface {
	CreateParkingLot() schema.CMDStore
	Park() schema.CMDStore
	Status() schema.CMDStore
	Help() schema.CMDStore
	ShellHistory() schema.CMDStore
	ParkHistory() schema.CMDStore
}
