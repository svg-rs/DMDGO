package chromedp

import (
	"context"
	"errors"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/svg-rs/DMDGO/utils"
	"github.com/urfave/cli/v3"
	"os"
	"time"
)

func Login(ctx context.Context, cmd *cli.Command) error {
	var err error
	var authHeader string
	threads := cmd.Int("threads")
	username := cmd.String("username")
	password := cmd.String("password")

	if username == "" || password == "" {
		return errors.New("username or password is empty")
	}

	for i := 0; i < int(threads); i++ {
		go func(i int) {

			var allocCtx context.Context
			var allocCancel context.CancelFunc
			var ctxChrome context.Context
			var cancel context.CancelFunc
			opts := append(chromedp.DefaultExecAllocatorOptions[:],
				chromedp.Flag("headless",
					true)) // Set headless to false for debugging

			allocCtx, allocCancel = chromedp.NewExecAllocator(ctx, opts...)
			defer allocCancel()
			ctxChrome, cancel = chromedp.NewContext(allocCtx)
			defer cancel()

			err = chromedp.Run(ctxChrome, network.Enable())
			if err != nil {
				utils.Error("Error enabling network monitoring:", err)
				return
			}

			err = chromedp.Run(ctxChrome, chromedp.ActionFunc(func(ctx context.Context) error {
				var cancelContext context.CancelFunc
				var listenCtx context.Context
				listenCtx, cancelContext = context.WithCancel(ctx)
				chromedp.ListenTarget(listenCtx, func(ev interface{}) {
					e, ok := ev.(*network.EventRequestWillBeSent)
					if ok {
						value, found := e.Request.Headers["Authorization"]
						if found {
							valStr, ok := value.(string)
							if ok {
								authHeader = valStr
								utils.Info(fmt.Sprintf("Authorization Header found in request to %s: %v", e.Request.URL, authHeader), nil)
								writeTokenToFile(authHeader)
								cancelContext()
							}
						}
					}
				})
				return nil
			}))
			if err != nil {
				utils.Error("Error setting up network listener:", err)
				return
			}

			err = chromedp.Run(ctxChrome,
				chromedp.Navigate("https://discord.com/login"),
				chromedp.WaitVisible(`input[name="email"]`, chromedp.ByQuery),
				chromedp.SendKeys(`input[name="email"]`, username, chromedp.ByQuery),
				chromedp.SendKeys(`input[name="password"]`, password, chromedp.ByQuery),
				chromedp.Click(`button[type="submit"]`, chromedp.ByQuery),
				chromedp.Sleep(5*time.Second),
			)
			if err != nil {
				utils.Error("Error logging in:", err)
				return
			}

			timeout := time.NewTimer(0 * time.Second)
			defer timeout.Stop()

			select {
			case <-timeout.C:
				if authHeader == "" {
					utils.Warn("Authorization Header not found within the time limit.", nil)
				} else {
					utils.Info(fmt.Sprintf("Authorization Header captured successfully. %s", authHeader), nil)
					utils.Info(fmt.Sprintf("Writing to file: ../output/tokens/logintokens.txt | %v ", authHeader), nil)
				}
			}
		}(i)
	}

	time.Sleep(15 * time.Second)
	return nil
}

func writeTokenToFile(writeString string) {
	if err := os.MkdirAll("output/tokens", os.ModePerm); err != nil {
		utils.Error("Error creating directory:", err)
		return
	}

	file, err := os.OpenFile("output/tokens/logintokens.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		utils.Error("Error opening file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(writeString + "\n")
	if err != nil {
		utils.Error("Error writing to file:", err)
	}
}
