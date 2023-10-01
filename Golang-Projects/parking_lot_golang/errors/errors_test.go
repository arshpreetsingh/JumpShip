package errors_test

import (
	"fmt"
	. "parking_lot/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Errors", func() {
	Context("Test ProcessFile", func() {
		It("ErrInvalidCommand", func() {
			res := ErrInvalidCommand("dummy")
			Expect(res.Error()).To(Equal(fmt.Sprintf(ErrInvalidCMD, "dummy")))
		})
		It("ErrInvalidArguments", func() {
			res := ErrInvalidArguments("park", 2, 1)
			Expect(res.Error()).To(Equal(fmt.Sprintf(ErrInsuffArguments, "park", 2, 1)))
		})
		It("ErrInvalidSlotCount", func() {
			res := ErrInvalidSlotCount(1)
			Expect(res.Error()).To(Equal(fmt.Sprintf(ErrInsuffSlot, 1)))
		})
		It("ErrNoCarFoundByColour", func() {
			res := ErrNoCarFoundByColour("red")
			Expect(res.Error()).To(Equal(fmt.Sprintf(NoCarFoundByColour, "red")))
		})
		It("ErrNoHistoyFound", func() {
			res := ErrNoHistoyFound("shell")
			Expect(res.Error()).To(Equal(fmt.Sprintf(NoHistoryFound, "shell")))
		})
		It("ErrDuplicateVehicle", func() {
			res := ErrDuplicateVehicle("ka-01-aj-1234")
			Expect(res.Error()).To(Equal(fmt.Sprintf(DuplicateVehicle, "ka-01-aj-1234")))
		})
	})
})
