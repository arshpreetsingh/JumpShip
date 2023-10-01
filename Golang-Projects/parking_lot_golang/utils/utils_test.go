package utils

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"parking_lot/errors"
)

var _ = Describe("Test utils functions:", func() {
	Context("Test SplitCmdArguments", func() {
		It("Should Fail - No tab space allowed", func() {
			input := "cmd \targ1 \t arg2"
			res, err := SplitCmdArguments(input)
			Expect(err).To(Equal(errors.ErrInvalidTabSpace))
			Expect(res).To(BeNil())
		})
		It("Should split command string", func() {
			input := "cmd arg1 arg2"
			res, err := SplitCmdArguments(input)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(len(res)).To(Equal(2))
		})
	})
	Context("Test IsRegNoValid", func() {
		It("Should Fail - Invalid Car Number", func() {
			input := "kaa-011-aj-12223"
			Expect(IsRegNoValid(input)).To(BeFalse())
		})
		It("Should Pass - Valid Car Number", func() {
			input := "TN-24-AJ-8462"
			Expect(IsRegNoValid(input)).To(BeTrue())
		})
	})
	Context("Test IsValidString", func() {
		It("Should Fail - Invalid string", func() {
			input := "balc!#-12223"
			Expect(IsValidString(input)).To(BeFalse())
		})
		It("Should Pass - Valid string", func() {
			input := "black"
			Expect(IsValidString(input)).To(BeTrue())
		})
	})

	It("Test time format", func() {
		fTime := FormatDateTime(time.Now())
		Expect(fTime).To(Equal(fTime))
	})
})
