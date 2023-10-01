package store

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"parking_lot/errors"
	"parking_lot/schema"
)

var _ = Describe("store packeg tests", func() {
	var (
		connection Store
	)
	connection = NewStore()

	It("Tear Down Store Data", func() {
		TearDown()
	})
	Context("Store connection tests", func() {
		TearDown()
		cmd := &schema.Command{
			Command: "help",
		}
		It("Help Store Execute", func() {
			res, err := connection.Help().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal(schema.AllCommendHint))
		})

		It("ParkHistory Store Execute, Without Parkinglot object", func() {
			res, err := connection.ParkHistory().Execute(cmd)
			Expect(err).To(Equal(errors.ErrNoParkingLot))
			Expect(res).To(Equal(""))
		})

		It("Shell Store Execute - without histories", func() {
			res, err := connection.ShellHistory().Execute(cmd)
			Expect(err).To(Equal(errors.ErrNoHistoyFound("shell")))
			Expect(res).To(Equal(""))
		})
		It("Shell Store Execute - with histories", func() {
			history := []*schema.IShellHistory{
				{
					Command:   "create_parking_lot 5",
					CreatedAt: time.Now(),
				},
				{
					Command:   "park KA-01-QW-1235 Red",
					CreatedAt: time.Now(),
				},
				{
					Command:   "leave 1",
					CreatedAt: time.Now(),
				},
			}
			cmd.RecordShellHistory(history)
			res, err := connection.ShellHistory().Execute(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res).To(Equal(res))
		})
	})
})
