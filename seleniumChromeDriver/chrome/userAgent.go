package chrome

type UserAgent string

const (
	Windows   UserAgent = "Windows"
	OSX       UserAgent = "OSX"
	OsxSafari UserAgent = "OsxSafari"

	IOS     UserAgent = "IOS"
	Android UserAgent = "Android"

	GoogleBot UserAgent = "GoogleBot"
)

var userAgents = map[UserAgent]string{
	Windows:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36",
	OSX:       "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko)",
	OsxSafari: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.2 Safari/605.1.15",
	IOS:       "Mozilla/5.0 (iPhone; CPU iPhone OS 13_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/77.0.3865.103 Mobile/15E148 Safari/605.1",
	GoogleBot: "Mozilla/5.0 (compatible; Googlebot/2.1; +http:// www.google.com / bot.html)",
	Android:   "Mozilla/5.0 (Linux; Android 8.0.0; SM-G960F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36",
}

var (
	chromeUserAgent = userAgents[Android]
)
