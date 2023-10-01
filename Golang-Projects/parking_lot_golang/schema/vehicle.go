package schema

import "strings"

type (
	//VehicleType holds the type of string
	VehicleType = string
)

const (
	// VehicleTypeTwoWheeler const vehicle type of `two_wheeler`
	VehicleTypeTwoWheeler VehicleType = "two_wheeler"
	// VehicleTypeCar const vehicle type of `car`
	VehicleTypeCar VehicleType = "car"
	// VehicleTypeBus const vehicle type of `bus`
	VehicleTypeBus VehicleType = "bus"
	// VehicleTypeTruck const vehicle type of `truck`
	VehicleTypeTruck VehicleType = "truck"
	// VehicleTypeAutoRickshow const vehicle type of `truck`
	VehicleTypeAutoRickshow VehicleType = "auto_rickshaw"
)

// Vehicle holds all the Vehicle properties
type Vehicle struct {
	RegistrationNumber string `json:"registration_type"`
	Colour             string `json:"colour"`
	Type               string `json:"type"`
	Model              string `json:"model"`
	Wheels             int    `json:"wheels"`
	Height             int    `json:"height"`
}

// GetRegNumber returns the vehicle reg number
func (v *Vehicle) GetRegNumber() string {
	return v.RegistrationNumber
}

// GetColour returns the vehicle color
func (v *Vehicle) GetColour() string {
	return v.Colour
}

// SetRegNumber sets the vehicle reg number
func (v *Vehicle) SetRegNumber(regNo string) {
	v.RegistrationNumber = regNo
	return
}

// SetColour sets the vehicle color
func (v *Vehicle) SetColour(colour string) {
	v.Colour = colour
	return
}

// IsVehicleColurMatched checks the colour match with the vehicle
func (v *Vehicle) IsVehicleColurMatched(colour string) bool {
	return (v.Colour == strings.ToLower(colour))
}

// IsVehicleRegNoMatched checks the reg number match with the vehicle
func (v *Vehicle) IsVehicleRegNoMatched(regNO string) bool {
	return (v.RegistrationNumber == regNO)
}
