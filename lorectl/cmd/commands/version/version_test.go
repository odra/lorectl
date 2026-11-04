package version

import (
	"fmt"
	"testing"
	"bytes"

	"github.com/spf13/cobra"

	. "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"

	"github.com/odra-ai-tools/lorectl/pkg/version"
)

var _ = Describe("root command", func() {
	var c *cobra.Command
 	var cmdOut bytes.Buffer
	var cmdErr bytes.Buffer

	BeforeEach(func() {
			cmdOut.Reset()
			cmdErr.Reset()
      c = Command()
			c.SetOut(&cmdOut)
			c.SetErr(&cmdErr)
	})

	It("Should print the program's version", func() {
		c.SetArgs([]string{})

		Expect(c.Execute()).To(Succeed())
		Expect(cmdOut.String()).To(Equal(fmt.Sprintf("v%s\n", version.Version)))
		Expect(cmdErr.String()).To(Equal(""))	
	})
})

func TestVersionCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VersionCmd Suite")
}
