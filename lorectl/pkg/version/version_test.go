package version

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"
)

var _ = Describe("root command", func() {

	It("should format version", func() {
		Expect(Version).To(Equal("0.0.0"))
	})
})

func TestVersion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Version Suite")
}
