package main

import (
	"log"
	"os"

	"github.com/bywachira/html-server/cli"
)

func main() {
	app := cli.SetupCLI()

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal("Error: ", err)
	}
}
