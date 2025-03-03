package selenium

import (
	"context"
	"github.com/svg-rs/DMDGO/utils"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"github.com/urfave/cli/v3"
	"time"
)

func Login(ctx context.Context, cmd *cli.Command) error {
	var err error
	var service *selenium.Service
	service, err = selenium.NewChromeDriverService("./chromedriver.exe", 4444)
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
	err = driver.Get("https://discord.com/login")
	if err != nil {
		utils.Error("Error getting page:", err)
	}
	var username string
	var password string
	username = cmd.String("username")
	password = cmd.String("password")
	usernameField, err := driver.FindElement(selenium.ByCSSSelector, "input[name='email']")
	passwordField, err := driver.FindElement(selenium.ByCSSSelector, "input[name='password']")
	err = usernameField.SendKeys(username)
	if err != nil {
		utils.Error("Error sending username:", err)
	}
	err = passwordField.SendKeys(password)
	if err != nil {
		utils.Error("Error sending password:", err)
	}
	loginButton, err := driver.FindElement(selenium.ByCSSSelector, "button[type='submit']")
	err = loginButton.Click()
	if err != nil {
		utils.Error("Error clicking login button:", err)
	}

	time.Sleep(5 * time.Second)
	// Wait for the page to load
	// edit code starting here

	time.Sleep(100000 * time.Second)
	return nil
}
