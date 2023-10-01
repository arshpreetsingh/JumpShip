package store

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("store packeg tests", func() {
	Context("Test Conenctions", func() {
		var s store
		s.CreateParkingLot()
		s.Park()
		s.Status()
		s.Help()
		s.ParkHistory()
		s.ShellHistory()
	})

	It("Tear Down Store Data", func() {
		TearDown()
	})
})

// Create store connection object
func Setup() {
}

func TearDown() {
	ParkingLot = nil
}
