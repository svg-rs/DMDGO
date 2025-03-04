package chromedp

import (
	"context"
	"github.com/svg-rs/DMDGO/utils"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"github.com/urfave/cli/v3"
)

func Test(ctx context.Context, cmd *cli.Command) error {
	var err error
	var service *selenium.Service
	service, err = selenium.NewChromeDriverService("./chromedriver.exe",
		4444)
	if err != nil {
		utils.Error("Error starting service:", err)
	}
	defer service.Stop()

	var caps selenium.Capabilities
	caps = selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chrome.Capabilities{Args: []string{}})

	var driver selenium.WebDriver
	driver, err = selenium.NewRemote(caps, "")
	if err != nil {
		utils.Error("Error creating driver:", err)
	}

	err = driver.Get("http://www.google.com")
	if err != nil {
		utils.Error("Error getting page:", err)
	}

	utils.Info("Selenium test completed successfully", nil)

	return nil
}
