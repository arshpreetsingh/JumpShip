package ishell

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("iShell Input Command File Process", func() {
	It("Should Fail - Invalid file", func() {
		fileName := "input.txt"
		ProcessFile(fileName)
	})
	Context("Test ProcessFile", func() {
		var tmpfile *os.File
		var err error
		var cmdInputStr = []byte(
			`create_parking_lot 6\n
        	park KA-01-HH-1234 White\n
        	park KA-01-HH-9999 White\n
        	park KA-01-BB-0001 Black\n
        	park KA-01-HH-7777 Red\n
        	park KA-01-HH-2701 Blue\n
        	park KA-01-HH-3141 Black\n
        	status\n
        	park KA-01-P-333 White\n
        	park DL-12-AA-9999 White\n
			exit\n`)
		BeforeEach(func() {
			InitIshell()
			tmpfile, err = ioutil.TempFile("", "file_inputs.txt")
			Ω(err).ShouldNot(HaveOccurred())
			_, err = tmpfile.Write(cmdInputStr)
			Ω(err).ShouldNot(HaveOccurred())
			err = tmpfile.Close()
			Ω(err).ShouldNot(HaveOccurred())
		})
		AfterEach(func() {
			Store = nil
		})

		It("Should split command string", func() {
			defer os.Remove(tmpfile.Name())
			ProcessFile(tmpfile.Name())
		})
	})
})
