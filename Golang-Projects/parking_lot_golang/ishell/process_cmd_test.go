package ishell

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"parking_lot/errors"
)

var _ = Describe("iShell Helpher functions", func() {
	Context("Test Process basic validations", func() {
		BeforeEach(func() {
			InitIshell()
		})
		AfterEach(func() {
			Store = nil
		})

		It("Invalid command to process", func() {
			availableCommandsHints()
			greeting()
			inputCmd := "invalid_command \targs"
			cmd, err := Process(inputCmd)
			Expect(err).To(Equal(errors.ErrInvalidTabSpace))
			Expect(cmd).To(BeNil())
		})
		It("Invalid command to process", func() {
			inputCmd := "invalid_command args"
			cmd, err := Process(inputCmd)
			Expect(err.Error()).To(Equal("Command 'invalid_command' is invalid!. Please type 'help' see available commands"))
			Expect(cmd).To(BeNil())
		})
		It("Should return valid cmd object for cmd create_parking_lot", func() {
			inputCmd := "create_parking_lot 6"
			cmd, err := Process(inputCmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(cmd.Command).To(Equal("create_parking_lot"))
			Expect(cmd.Arguments).To(Equal([]string{"6"}))
			Expect(cmd.ShellHistory).To(BeNil())
		})
		It("Should return valid cmd object for cmd park", func() {
			inputCmd := "park KA-01-HH-1234 White"
			cmd, err := Process(inputCmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(cmd.Command).To(Equal("park"))
			Expect(cmd.Arguments).To(Equal([]string{"KA-01-HH-1234", "White"}))
			Expect(cmd.ShellHistory).To(BeNil())
		})
		It("Should return valid cmd object for cmd help", func() {
			inputCmd := "help"
			cmd, err := Process(inputCmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(cmd.Command).To(Equal("help"))
			Expect(len(cmd.Arguments)).To(Equal(0))
			Expect(cmd.ShellHistory).To(BeNil())
		})
		It("Should return valid cmd object for cmd shell_history", func() {
			inputCmd := "shell_history"
			cmd, err := Process(inputCmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(cmd.Command).To(Equal("shell_history"))
			Expect(len(cmd.Arguments)).To(Equal(0))
			Expect(cmd.ShellHistory).To(BeNil())
		})
		It("Should return valid cmd object for cmd park_history", func() {
			inputCmd := "park_history"
			cmd, err := Process(inputCmd)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(cmd.Command).To(Equal("park_history"))
			Expect(len(cmd.Arguments)).To(Equal(0))
			Expect(cmd.ShellHistory).To(BeNil())
		})

	})
})
