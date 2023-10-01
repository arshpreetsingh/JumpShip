package schema_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"parking_lot/errors"
	. "parking_lot/schema"
)

var _ = Describe("models commands", func() {
	Context("Test ProcessFile", func() {
		var cmd *Command
		var ishell *IShell
		BeforeEach(func() {
			cmd = &Command{
				Command:   CMDPark,
				Arguments: []string{"KA-01-AJ-1234", "White"},
			}

			ishell = &IShell{
				Prompt: DefaultPrompt,
			}
		})
		AfterEach(func() {
			cmd = nil
			ishell = nil
		})

		It("ShowPrompt", func() {
			ishell.ShowPrompt()
		})
		It("RecordHistory", func() {
			Expect(len(ishell.History)).To(Equal(0))
			ishell.RecordHistory("red")
			Expect(len(ishell.History)).To(Equal(1))
			cmd.RecordShellHistory(ishell.History)
			Expect(len(cmd.ShellHistory)).To(Equal(1))
		})
	})
	Context("Parking Lot Model Helper function", func() {
		var pl *ParkingLot
		It("FirstAvailableSlot", func() {
			pl = new(ParkingLot)
			pl.Name = "test_parking_lot"
			pl.TotalSlots = 5
			pl.Slots = make([]*Slot, 5)
			for i := range pl.Slots {
				pl.Slots[i] = new(Slot)
				pl.Slots[i].SetID(i + 1)
				pl.Slots[i].SetName(i + 1)
				pl.Slots[i].MakeSlotFree()
			}
			slot, err := pl.FirstAvailableSlot()
			if err != nil {
				fmt.Println(err)
			}
			//N)(err).ShouldNot(HaveOccurred())
			Expect(slot.GetID()).To(Equal(uint(1)))
			Expect(slot.GetName()).To(Equal("slot_1"))
			Expect(slot.IsFree).To(BeTrue())
			Expect(slot.Vehicle).To(BeNil())
			Expect(slot.IsSlotOccupied()).To(BeFalse())
		})
		It("FirstAvailableSlot", func() {
			pl = new(ParkingLot)
			slot, err := pl.FirstAvailableSlot()
			Expect(err).To(Equal(errors.ErrParkingSlotsFull))
			Expect(slot).To(BeNil())
		})
	})
	Context("Slot Model Helper function", func() {
		slot := new(Slot)
		slot.SetID(1)
		slot.SetName(1)

		It("Park Vehicle in slot", func() {
			v := new(Vehicle)
			v.SetColour("Red")
			v.SetRegNumber("TN-24-AJ-8462")
			err := slot.ParkVehicle(v)
			fmt.Println(err)
			//N)(err).ShouldNot(HaveOccurred())
			Expect(slot.IsFree).To(BeFalse())
			Expect(slot.Vehicle.GetColour()).To(Equal("Red"))
			Expect(slot.Vehicle.GetRegNumber()).To(Equal("TN-24-AJ-8462"))
		})
		It("Park Vehicle in slot fail", func() {
			v := new(Vehicle)
			v.SetColour("Red")
			v.SetRegNumber("TN-24-AJ-8462")
			err := slot.ParkVehicle(v)
			Expect(err).To(Equal(errors.ErrSlotAlreadyOccupied))
			Expect(slot.IsFree).To(BeFalse())
		})
		It("Leave from slot okay", func() {
			err := slot.ExitPark()
			fmt.Println(err)
			//N)(err).ShouldNot(HaveOccurred())
			Expect(slot.IsFree).To(BeTrue())
			Expect(slot.Vehicle).To(BeNil())
		})
		It("Leave from slot fail", func() {
			err := slot.ExitPark()
			Expect(err).To(Equal(errors.ErrSlotAlreadyAvailable))
			Expect(slot.IsFree).To(BeTrue())
			Expect(slot.Vehicle).To(BeNil())
		})
	})
	Context("Vehicle Model Helper function", func() {
		v := new(Vehicle)
		v.SetColour("red")
		v.SetRegNumber("TN-24-AJ-8462")
		It("Helpher validators", func() {
			Expect(v.IsVehicleColurMatched("Red")).To(BeTrue())
			Expect(v.IsVehicleColurMatched("Black")).To(BeFalse())
			Expect(v.IsVehicleRegNoMatched("TN-24-AJ-8462")).To(BeTrue())
			Expect(v.IsVehicleRegNoMatched("TN-24-AJ-84623")).To(BeFalse())
		})
	})
})
