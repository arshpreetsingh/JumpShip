package store

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"parking_lot/errors"
	"parking_lot/schema"
)

var _ = Describe("parking lot store tests", func() {
	var (
		connection Store
	)
	connection = NewStore()
	It("Tear Down Store Data", func() {
		TearDown()
	})
	Context("parking_lot store excute", func() {
		TearDown()
		cmd := &schema.Command{
			Command: "create_parking_lot",
		}
		It("create_parking_lot help", func() {
			cmd.Arguments = []string{"help"}
			res, err := connection.CreateParkingLot().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal(schema.CMDCreateParkingLotHint))
		})

		It("invalid arguments string", func() {
			cmd.Arguments = []string{"assa"}
			res, err := connection.CreateParkingLot().Execute(cmd)
			Expect(err).To(Equal(errors.ErrInvalidInputSlot))
			Expect(res).To(Equal(""))
		})

		It("invalid arguments negative slot count", func() {
			cmd.Arguments = []string{"-10"}
			res, err := connection.CreateParkingLot().Execute(cmd)
			Expect(err).To(Equal(errors.ErrInvalidSlotCount(-10)))
			Expect(res).To(Equal(""))
		})

		It("Invalid total slot count", func() {
			cmd.Arguments = []string{"0"}
			res, err := connection.CreateParkingLot().Execute(cmd)
			Expect(err).To(Equal(errors.ErrInvalidSlotCount(0)))
			Expect(res).To(Equal(""))
		})

		It("Create a parking lot with 5 slots", func() {
			cmd.Arguments = []string{"5"}
			res, err := connection.CreateParkingLot().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal(fmt.Sprintf(ParkinglotCreatedInfo, 5)))
		})

		It("Create a parking lot with 100 slots", func() {
			TearDown()
			cmd.Arguments = []string{"100"}
			res, err := connection.CreateParkingLot().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal(fmt.Sprintf(ParkinglotCreatedInfo, 100)))
		})

		It("parking lot already created", func() {
			cmd.Arguments = []string{"5"}
			res, err := connection.CreateParkingLot().Execute(cmd)
			Expect(err).To(Equal(errors.ErrParkingLotAlreadyCreated))
			Expect(res).To(Equal(""))
		})
	})
})
