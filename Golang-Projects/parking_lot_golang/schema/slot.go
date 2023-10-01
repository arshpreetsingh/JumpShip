package schema

import (
	"fmt"
	"parking_lot/errors"
)

// Slot holds all the slot properties
type Slot struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
	IsFree    bool     `json:"is_free"`
	BlockName string   `json:"block_name"`
	BlockID   uint     `json:"block_id"`
	Vehicle   *Vehicle `json:"vehicle"`
}

// GetID returns slot id
func (s *Slot) GetID() uint {
	return s.ID
}

// GetName returns slot name
func (s *Slot) GetName() string {
	return s.Name
}

// SetID assign the id to slot
func (s *Slot) SetID(ID int) {
	s.ID = uint(ID)
	return
}

// MakeSlotFree updates the slot as free
func (s *Slot) MakeSlotFree() {
	s.IsFree = true
	return
}

// SetSlotOccupied updates the slot as occupied
func (s *Slot) SetSlotOccupied() {
	s.IsFree = false
	return
}

// SetName assigns the name to slot
func (s *Slot) SetName(ID int) {
	s.Name = fmt.Sprintf("slot_%d", ID)
	return
}

// IsSlotAvailable checks if the slot is occupied or not
func (s *Slot) IsSlotAvailable() bool {
	return (s.IsFree && s.Vehicle == nil)
}

// IsSlotOccupied checks if the slot is occupied or not
func (s *Slot) IsSlotOccupied() bool {
	return (s.Vehicle != nil)
}

// ParkVehicle assign/park the Vehicle in the slot
func (s *Slot) ParkVehicle(car *Vehicle) error {
	if s.Vehicle != nil {
		return errors.ErrSlotAlreadyOccupied
	}
	// park vehicle here, make slot occupied
	s.SetSlotOccupied()
	s.Vehicle = car
	return nil
}

// ExitPark makes the slot available to park next Vehicle
func (s *Slot) ExitPark() error {
	if s.Vehicle == nil {
		return errors.ErrSlotAlreadyAvailable
	}
	s.MakeSlotFree()
	s.Vehicle = nil
	return nil
}
