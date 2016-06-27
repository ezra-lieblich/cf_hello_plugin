package main
import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
)
type HelloPlugin struct{}

func (c *HelloPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	cliConnection.CliCommand()
	if len(args)== 1{
		fmt.Println("Hello World")
	}else{
		fmt.Print("Hello")
		for _, value := range args[1:]{
			fmt.Print(" " + value)
		}
		fmt.Println("")
	}

}

func (c *HelloPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "HelloWorldPlugin",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 4,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name: "hello-world",
				HelpText: "Prints the Hello + user",
				UsageDetails: plugin.Usage{
					Usage: "hello-world\n	cf hello-world text",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(HelloPlugin))
}