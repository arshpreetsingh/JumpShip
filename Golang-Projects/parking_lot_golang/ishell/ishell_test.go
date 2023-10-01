package ishell_test

import (
	"io/ioutil"
	"os"

	. "parking_lot/ishell"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("iShell Input Command Process", func() {
	It("Invalid command", func() {
		cmd := []byte("invalid_cmd 6\n")
		tmpfile, err := ioutil.TempFile("", "example")
		Ω(err).ShouldNot(HaveOccurred())

		defer os.Remove(tmpfile.Name()) // clean up
		_, err = tmpfile.Write(cmd)
		Ω(err).ShouldNot(HaveOccurred())
		_, err = tmpfile.Seek(0, 0)
		Ω(err).ShouldNot(HaveOccurred())

		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }() // Restore original Stdin

		os.Stdin = tmpfile
		Start()

		err = tmpfile.Close()
		Ω(err).ShouldNot(HaveOccurred())
	})

	Context("Interactive shell - test commands", func() {
		BeforeEach(func() {
			InitIshell()
		})
		AfterEach(func() {
			Store = nil
		})

		It("Valid command - should log the stdout", func() {
			cmd := []byte("create_parking_lot 6\n")
			tmpfile, err := ioutil.TempFile("", "example")
			Ω(err).ShouldNot(HaveOccurred())

			defer os.Remove(tmpfile.Name()) // clean up
			_, err = tmpfile.Write(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			_, err = tmpfile.Seek(0, 0)
			Ω(err).ShouldNot(HaveOccurred())

			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }() // Restore original Stdin

			os.Stdin = tmpfile
			Start()

			err = tmpfile.Close()
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Exit command", func() {
			cmd := []byte("exit\n")
			tmpfile, err := ioutil.TempFile("", "example")
			Ω(err).ShouldNot(HaveOccurred())

			defer os.Remove(tmpfile.Name()) // clean up
			_, err = tmpfile.Write(cmd)
			Ω(err).ShouldNot(HaveOccurred())
			_, err = tmpfile.Seek(0, 0)
			Ω(err).ShouldNot(HaveOccurred())

			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }() // Restore original Stdin

			os.Stdin = tmpfile
			Start()

			err = tmpfile.Close()
			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})
