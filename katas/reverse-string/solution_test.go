package kata_test

import (
	"testing"

	. "github.com/jfkonecn/code-wars-katas/katas/reverse-string"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSolution(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Solution Suite")
}

var _ = Describe("Test Example", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(Solution("world")).To(Equal("dlrow"))
	})
})
