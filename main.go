package main

import (
	"context"
	"github.com/fatih/color"
	"os"

	"github.com/svg-rs/DMDGO/selenium"
	"github.com/svg-rs/DMDGO/utils"
	"github.com/urfave/cli/v3"
)

func main() {
	var cmd *cli.Command
	cmd = setupCLI()

	var err error
	err = cmd.Run(context.Background(), os.Args)
	if err != nil {
		utils.Error("Application error:", err)
	}
}

func setupCLI() *cli.Command {
	var cmd *cli.Command
	cmd = &cli.Command{
		Name:  color.New(color.FgBlue, color.Bold).Sprintf("DMDGO"),
		Usage: color.New(color.FgYellow).Sprintf("A rewrite of DMDGO"),
		Commands: []*cli.Command{
			{
				Name:   "test",
				Usage:  color.New(color.FgYellow).Sprintf("Test to see if selenium is working"),
				Action: selenium.Test,
			},
		},
	}
	return cmd
}
