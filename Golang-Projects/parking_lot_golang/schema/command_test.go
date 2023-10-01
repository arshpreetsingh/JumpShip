package schema_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"parking_lot/errors"
	. "parking_lot/schema"
)

var _ = Describe("Command", func() {
	Context("Test ProcessFile", func() {
		var cmd *Command
		BeforeEach(func() {
			cmd = &Command{
				Command:   "park",
				Arguments: []string{"KA-01-AJ-1234", "White"},
			}
		})
		AfterEach(func() {
			cmd = nil
		})

		It("GetName", func() {
			Expect(cmd.GetName()).To(Equal("park"))
		})
		It("GetArguments", func() {
			Expect(len(cmd.GetArguments())).To(Equal(2))
		})
		It("IsExit", func() {
			Expect(cmd.IsExit()).To(BeFalse())
		})
		It("Validate Invalid Command", func() {
			cmd = &Command{
				Command: "invalid",
			}
			err := cmd.Ok()
			Expect(err).To(Equal(errors.ErrInvalidCommand(cmd.Command)))
		})
		It("Validate Invalid Argument", func() {
			cmd.Arguments = []string{}
			err := cmd.Ok()
			Expect(err).To(Equal(errors.ErrInvalidArguments(cmd.Command, CMDArgumentLength[cmd.Command], len(cmd.Arguments))))
		})
		It("Validate pass", func() {
			err := cmd.Ok()
			Expect(err).To(BeNil())
		})
	})
})
