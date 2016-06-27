package main
import (
	"github.com/cloudfoundry/cli/plugin/pluginfakes"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"CLI-Hello/git-files/testhelpers/io"
)
var _ = ginkgo.Describe("HelloWorld", func() {
	var(
		connection    *pluginfakes.FakeCliConnection
		ExamplePlugin             *HelloPlugin
	)
	ginkgo.BeforeEach(func() {
		connection = &pluginfakes.FakeCliConnection{}
		ExamplePlugin = &HelloPlugin{}
	})
	ginkgo.Context("no parameters", func() {
		ginkgo.It("displays the default Hello World", func(){
			output := io.CaptureOutput(func(){
				ExamplePlugin.Run(connection, []string{"hello-world"})
			})
			gomega.Expect(output).To(gomega.Equal([]string{"Hello World", ""}))
		})
	})
	ginkgo.Context("more than one parameter", func(){
		ginkgo.It("displays one name", func(){
			output := io.CaptureOutput(func(){
				ExamplePlugin.Run(connection, []string{"hello-world", "ezra"})
			})
			gomega.Expect(output).To(gomega.Equal([]string{"Hello ezra", ""}))
		})
		ginkgo.It("displays a full name", func(){
			output := io.CaptureOutput(func(){
				ExamplePlugin.Run(connection, []string{"hello-world", "ezra",
					"marcus", "lieblich"})
			})
			gomega.Expect(output).To(gomega.Equal([]string{"Hello ezra marcus lieblich", ""}))

		})
	})

})