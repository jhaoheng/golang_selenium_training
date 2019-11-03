package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/firefox"
)

func main() {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// These paths will be different on your system.
		seleniumPath    = "./vendor/selenium-server-standalone-4.0.0-alpha-1.jar"
		geckoDriverPath = "./vendor/geckodriver/geckodriver-v0.26.0-macos"
		// geckoDriverPath = "./vendor/geckodriver/geckodriver-v0.26.0-linux64"
		port = 8080
	)
	opts := []selenium.ServiceOption{
		// selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{
		"browserName": "firefox",
		"version":     "51",
		// "platform":     "LINUX",
		"maxInstances": 1,
		"marionette":   true,
		"cleanSession": true,
	}

	firefoxCaps := firefox.Capabilities{
		// Binary: "",
		Args: []string{
			// "--disable-logging",
			// "--disable-extensions",
			// "--headless", // 設定Chrome無頭模式，在linux下執行，需要設定這個引數，否則會報錯
			// "--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", // 模擬user-agent，防反爬
		},
	}
	caps.AddFirefox(firefoxCaps)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("http://play.golang.org/?simple=1"); err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#code")
	if err != nil {
		panic(err)
	}
	// Remove the boilerplate code already in the text box.
	if err := elem.Clear(); err != nil {
		panic(err)
	}

	fmt.Println(elem)
	var stop1 bool
	fmt.Scanln(&stop1)

	// Enter some new code in text box.
	err = elem.SendKeys(`
		package main
		import "fmt"

		func main() {
			fmt.Println("Hello WebDriver!\n")
		}
	`)
	if err != nil {
		panic(err)
	}

	// Click the run button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}

	// Wait for the program to finish running and get the output.
	outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "#output")
	if err != nil {
		panic(err)
	}

	var output string
	for {
		output, err = outputDiv.Text()
		if err != nil {
			panic(err)
		}
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))

	var stop bool
	fmt.Scanln(&stop)
}
