package main

import (
	"log"
	"os"

	"github.com/takutakahashi/snip/cmd/operation"
	"github.com/urfave/cli/v2"
)

func main() {
	export := operation.CommandExport()
	app := &cli.App{
		Name:   "snip",
		Action: export.Action,
		Flags:  export.Flags,
		Commands: []*cli.Command{
			operation.CommandNew(),
			operation.CommandExec(),
			operation.CommandList(),
			operation.CommandShow(),
			operation.CommandUpdate(),
			operation.CommandInit(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
