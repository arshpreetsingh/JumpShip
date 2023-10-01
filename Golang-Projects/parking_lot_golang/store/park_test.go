package store

import (
	"fmt"

	"parking_lot/errors"
	"parking_lot/schema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("parking lot store tests", func() {
	var (
		connection Store
	)
	connection = NewStore()
	It("Tear Down Store Data", func() {
		TearDown()
	})

	Context("parking_lot store execute", func() {
		TearDown()

		It("park help", func() {
			cmd := &schema.Command{
				Command:   "park",
				Arguments: []string{"help"},
			}
			res, err := connection.Park().Execute(cmd)
			N)(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal(schema.CMDParkHint))
		})
		It("No parking lot available", func() {
			cmd := &schema.Command{
				Command:   "park",
				Arguments: []string{"ka-02-aw-1234", "123"},
			}
			res, err := connection.Park().Execute(cmd)

			Expect(err).To(Equal(errors.ErrNoParkingLot))
			Expect(res).To(Equal(""))
		})

		It("Create a parking lot with 5 slots", func() {
			cmd := &schema.Command{
				Command:   "create_parking_lot",
				Arguments: []string{"1"},
			}
			res, err := connection.CreateParkingLot().Execute(cmd)
			N)(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal(fmt.Sprintf(ParkinglotCreatedInfo, 1)))
		})

		It("invalid arguments registration number", func() {
			cmd := &schema.Command{
				Command:   "park",
				Arguments: []string{"k2a-123-21-s123", "Red"},
			}
			res, err := connection.Park().Execute(cmd)
			Expect(err).To(Equal(errors.ErrInvalidRegNo))
			Expect(res).To(Equal(""))
		})
		It("invalid arguments colour", func() {
			cmd := &schema.Command{
				Command:   "park",
				Arguments: []string{"ka-02-aw-1234", "123!@"},
			}
			res, err := connection.Park().Execute(cmd)
			Expect(err).To(Equal(errors.ErrInvalidColour))
			Expect(res).To(Equal(""))
		})

		It("no park history", func() {
			cmd := &schema.Command{
				Command:   "park_history",
				Arguments: []string{},
			}
			res, err := connection.ParkHistory().Execute(cmd)
			Expect(err).To(Equal(errors.ErrNoHistoyFound("parking")))
			Expect(res).To(Equal(""))
		})

		It("park a vehicle", func() {
			cmd := &schema.Command{
				Command:   "park",
				Arguments: []string{"TN-24-AJ-8462", "Red"},
			}
			res, err := connection.Park().Execute(cmd)
			N)(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal("Allocated slot number: 1"))
		})
		It("park a vehicle - already slots full", func() {
			cmd := &schema.Command{
				Command:   "park",
				Arguments: []string{"TN-24-AJ-8442", "Red"},
			}
			res, err := connection.Park().Execute(cmd)
			Expect(err).To(Equal(errors.ErrParkingSlotsFull))
			Expect(res).To(Equal(""))
		})
		It("park history", func() {
			cmd := &schema.Command{
				Command:   "park_history",
				Arguments: []string{},
			}
			res, err := connection.ParkHistory().Execute(cmd)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(res))
		})
		It("car already parked", func() {
			cmd := &schema.Command{
				Command:   "park",
				Arguments: []string{"TN-24-AJ-8462", "Red"},
			}
			res, err := connection.Park().Execute(cmd)
			Expect(err).To(Equal(errors.ErrDuplicateVehicle(cmd.Arguments[0])))
			Expect(res).To(Equal(""))
		})
	})
})
