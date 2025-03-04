package main

import (
	"context"

	"github.com/fatih/color"
	"github.com/svg-rs/DMDGO/chromedp"
	"github.com/svg-rs/DMDGO/utils"
	"github.com/urfave/cli/v3"
	"os"
)

func main() {
	cmd := setupCLI()
	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		utils.Error("Application error:", err)
	}
}

func setupCLI() *cli.Command {
	cmd := &cli.Command{
		Name:  color.New(color.FgBlue, color.Bold).Sprintf("DMDGO"),
		Usage: color.New(color.FgYellow).Sprintf("A rewrite of DMDGO in Selenium"),
		Commands: []*cli.Command{
			{
				Name:  "test",
				Usage: color.New(color.FgYellow).Sprintf("Test to see if chromedp is working"),
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "threads",
						Aliases: []string{"t"},
						Usage:   "Number of goroutines to run concurrently",
						Value:   1,
					},
				},
				Action: chromedp.Test,
			},
			{
				Name: "login",
				Usage: color.New(color.FgHiCyan).Sprintf(
					"Login to Discord | -u <username> -p <password> -t <threads> -uf <usernamefile> -pf <passwordfile>"),
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
					&cli.StringFlag{
						Name:     "usernamefile",
						Aliases:  []string{"uf"},
						Usage:    "File containing Discord usernames, one per line",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "passwordfile",
						Aliases:  []string{"pf"},
						Usage:    "File containing Discord passwords, one per line",
						Required: false,
					},
					&cli.IntFlag{
						Name:    "threads",
						Aliases: []string{"t"},
						Usage:   "Number of goroutines to run concurrently",
						Value:   1,
					},
				},
				Action: chromedp.Login,
			},
		},
	}
	return cmd
}
