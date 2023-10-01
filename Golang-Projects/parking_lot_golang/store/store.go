package store

import "parking_lot/schema"

var (
	// ParkinglotCreatedInfo holds the STDOUT message for cmd `create_parking_lot`
	ParkinglotCreatedInfo = "Created​ a parking​ lot with %d slots"
	// SlotAllocatedInfo holds the STDOUT message for cmd `park`
	SlotAllocatedInfo = "Allocated slot number: %v"
	// SlotIsFreeInfo holds the STDOUT message for cmd `status`
	SlotIsFreeInfo = "Slot number %v is free"
)

// ParkingLot holds the all parking data
var ParkingLot *schema.ParkingLot

type store struct {
	createParkingLot schema.CMDStore
	park             schema.CMDStore
	status           schema.CMDStore
	help             schema.CMDStore
	shellHistory     schema.CMDStore
	parkHistory      schema.CMDStore
}

func (s store) CreateParkingLot() schema.CMDStore {
	return s.createParkingLot
}

func (s store) Park() schema.CMDStore {
	return s.park
}

func (s store) Status() schema.CMDStore {
	return s.status
}

func (s store) Help() schema.CMDStore {
	return s.help
}

func (s store) ShellHistory() schema.CMDStore {
	return s.shellHistory
}

func (s store) ParkHistory() schema.CMDStore {
	return s.parkHistory
}

// NewStore returns the store object
func NewStore() *store {
	st := InitStore()
	st.createParkingLot = NewCreateParkingLotStore(st)
	st.park = NewParkStore(st)
	st.status = NewStatusStore(st)

	st.help = NewHelpStore(st)
	st.shellHistory = NewShellHistoryStore(st)
	st.parkHistory = NewParkHistoryStore(st)

	return st
}

// InitStore returns the store object
func InitStore() *store {
	return new(store)
}
