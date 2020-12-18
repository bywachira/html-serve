package cli

import (
	"fmt"
	"sort"

	"github.com/bywachira/html-server/server"

	"github.com/urfave/cli/v2"
)

// SetupCLI initialize cli
func SetupCLI() *cli.App {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "Serve your HTML file",
				Action: func(c *cli.Context) error {
					fmt.Println("We are serving this directory")

					server.Serve()
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	return app
}
