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
		Name: color.New(color.FgBlue, color.Bold).Sprintf("DMDGO"),
		Usage: color.New(color.FgYellow).Sprintf(
			"A rewrite of DMDGO in Selenium"),
		Commands: []*cli.Command{
			{
				Name:   "test",
				Usage:  color.New(color.FgYellow).Sprintf("Test to see if selenium is working"),
				Action: selenium.Test,
			},
			{
				Name:  "login",
				Usage: color.New(color.FgHiCyan).Sprintf("Login to Discord"),
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "username",
						Aliases:  []string{"u"},
						Usage:    "Discord username",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "password",
						Aliases:  []string{"p"},
						Usage:    "Discord password",
						Required: true,
					},
				},
				Action: selenium.Login,
			},
		},
	}
	return cmd
}
