package main

import (
	"fmt"
	"seleniumChromeDriver/chrome"

	"github.com/fatih/color"
)

func main() {
	RunWithChromeWindow := false
	cObj := chrome.NewAgent(RunWithChromeWindow)
	cObj.RunWebDriver()
	webDriver := cObj.GetWebDriver()
	defer cObj.CloseAgent()

	webDriver.Get("http://google.com")

	red := color.New(color.FgRed, color.Bold)
	var stop bool
	red.Printf("\n\nPlease Press Enter To Close ...")
	fmt.Scanln(&stop)
}
