package root

import (
	"testing"

	"github.com/spf13/cobra"

	. "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"
)

var _ = Describe("root command", func() {
	var c *cobra.Command

	BeforeEach(func() {
      c = Command()
	})

	It("should use defaults", func() {
		c.SetArgs([]string{})
		Expect(c.Execute()).To(Succeed())

		verbose, _ := c.Flags().GetBool("verbose")		
		Expect(verbose).To(BeFalse())

		config, _ := c.Flags().GetString("config")		
		Expect(config).To(Equal(""))
	})

	It("should not use defaults", func() {
		c.SetArgs([]string{"--verbose", "--config", "/etc/lorectl.yml"})
		Expect(c.Execute()).To(Succeed())

		verbose, _ := c.Flags().GetBool("verbose")		
		Expect(verbose).To(BeTrue())

		config, _ := c.Flags().GetString("config")		
		Expect(config).To(Equal("/etc/lorectl.yml"))
	})
})

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Main Suite")
}
/*
Expect(c.ParseFlags([]string{"--verbose"})).To(Succeed())
        v, _ := c.Flags().GetBool("verbose")
        Expect(v).To(BeTrue())
*/
