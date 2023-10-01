package ishell_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIshell(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ishell Suite")
}
