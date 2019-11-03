package chrome

/*
How to use

````
cObj := chrome.NewAgent()
cObj.RunWebDriver() or cObj.RunWebDriverByProxy("YOUR PROXY SERVER IP", PORT)
defer cObj.CloseAgent()
webDriver := cObj.GetWebDriver()
````
*/

import (
	"fmt"
	"os"
	"runtime"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"

	slog "github.com/tebeka/selenium/log"

	prettyJsonLog "github.com/happierall/l"
)

const (
	port = 8080

	LINUX   string = "linux"
	MAC_OSX string = "darwin"

	LinuxChromeDriverPath  string = "./bin/chromedriver/chromedriver-linux64"
	DarwinChromeDriverPath string = "./bin/chromedriver/chromedriver-darwin"
)

type IChromeDriver interface {
	RunWebDriver()
	RunWebDriverByProxy(ip string, port int)
	CloseAgent()
	GetWebDriver() selenium.WebDriver
	ShowChromeDriverPath()
	ShwoSeleniumCaps()
}

type ChromeObj struct {
	binPath   string
	caps      selenium.Capabilities
	Browser   *selenium.Service
	WebDriver selenium.WebDriver
}

func NewAgent(RunWithChromeWindow bool) IChromeDriver {
	cObj := &ChromeObj{}

	// 設定 chromeDriver path
	cObj.setChrmoeDriverPath()
	// 設定 Selenium Capabilities
	cObj.setSeleniumCapabilities(RunWithChromeWindow)
	// 設定 browser
	cObj.setBrowser()
	return cObj
}

func (cObj *ChromeObj) RunWebDriverByProxy(proxyIp string, proxyPort int) {
	// proxy
	proxy := selenium.Proxy{
		Type:     selenium.Manual,
		HTTP:     proxyIp,
		HTTPPort: proxyPort,
	}
	cObj.caps.AddProxy(proxy)
	cObj.buildWebDriver()
}

func (cObj *ChromeObj) RunWebDriver() {
	cObj.buildWebDriver()
}

func (cObj *ChromeObj) CloseAgent() {
	cObj.CloseBrowser()
	cObj.QuitWebDriver()
}

func (cObj *ChromeObj) setChrmoeDriverPath() {
	if runtime.GOOS == MAC_OSX {
		cObj.binPath = DarwinChromeDriverPath
	} else if runtime.GOOS == LINUX {
		cObj.binPath = LinuxChromeDriverPath
	} else {
		panic("There have no 'chrmoedriver' for this OS")
	}
}

func (cObj *ChromeObj) setSeleniumCapabilities(RunWithChromeWindow bool) {

	args := []string{
		// "ignore-certificate-errors",
		"--disable-crash-reporter",
		"--disable-demo-mode",
		"--disable-cookie-encryption",
		"--disable-component-cloud-policy",
		"--disable-checker-imaging",
		"--disable-bundled-ppapi-flash", // 禁止 flash
		"--disable-internal-flash",      // 禁止 flash
		// "--disable-prompt-on-repost",
		"--disable-logging",
		"--log-level=3",
		"--disable-extensions",
		"--no-sandbox",
		"--user-agent=" + chromeUserAgent, // 模擬user-agent
	}
	if !RunWithChromeWindow {
		args = append(args, "--headless") // 設定Chrome無頭模式，在linux下執行，需要設定這個引數，否則會報錯
	}

	// 設定 Capabilities
	caps := selenium.Capabilities{
		"browserName": "chrome",
		"Platform":    "Linux",
	}
	// ChromeDriver ref : `https://sites.google.com/a/chromium.org/chromedriver/capabilities`
	// Args ref :`https://peter.sh/experiments/chromium-command-line-switches/`
	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: args,
		Prefs: map[string]interface{}{
			"profile.managed_default_content_settings.images": 2, // 禁止加載圖片
			// {"--disable-bundled-ppapi-flash": 2},              //禁止FLSH
		},
	}
	caps.AddChrome(chromeCaps)

	//
	caps.SetLogLevel(slog.Server, slog.Off)
	caps.SetLogLevel(slog.Browser, slog.Off)
	caps.SetLogLevel(slog.Client, slog.Off)
	caps.SetLogLevel(slog.Driver, slog.Off)
	caps.SetLogLevel(slog.Performance, slog.Off)
	caps.SetLogLevel(slog.Profiler, slog.Off)

	(*cObj).caps = caps
}

func (cObj *ChromeObj) setBrowser() {
	opts := []selenium.ServiceOption{
		// Enable fake XWindow session.
		// selenium.StartFrameBuffer(),
		selenium.Output(os.Stderr), // Output debug information to STDERR
	}

	// Enable debug info.
	selenium.SetDebug(false)

	// Starts a ChromeDriver instance in the background. (This is browser)
	ChromeDriverService, err := selenium.NewChromeDriverService((*cObj).binPath, port, opts...)
	if err != nil {
		panic(err)
	}
	(*cObj).Browser = ChromeDriverService
}

/*
WebDriver
*/
func (cObj *ChromeObj) buildWebDriver() {
	webDriver, err := selenium.NewRemote((*cObj).caps, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	(*cObj).WebDriver = webDriver
}

// usage : defer ChromeObj.CloseBrowser()
func (cObj *ChromeObj) CloseBrowser() {
	(*cObj).Browser.Stop()
}

// usage : defer ChromeObj.QuitWebDriver()
func (cObj *ChromeObj) QuitWebDriver() {
	cObj.WebDriver.Quit()
}

// Show func
func (CObj *ChromeObj) ShowChromeDriverPath() {
	fmt.Printf("The ChromeDriver Path is : %s\n", CObj.binPath)
}

func (CObj *ChromeObj) ShwoSeleniumCaps() {
	prettyJsonLog.Print(CObj.caps)
}

func (cObj *ChromeObj) GetWebDriver() selenium.WebDriver {
	return cObj.WebDriver
}
