# How to use

## SeleniumChromeDriver
- 非 Docker 環境, 直接執行 `go run main.go`(須確認本機有安裝 chrome browser)
- Docker 環境
    1. 在 main.go 中, 將 `RunWithChromeWindow = false`
    2. `docker-compose up -d` and `docker exec -it app /bin/bash`
    3. go run main.go

# QA
- selenium - chromedriver : 如何關掉 devtool?
    - 找不到方法關掉 - 20191025 
- 如何下載 bin:chromedriver
    1. 載點 : https://chromedriver.chromium.org/downloads
    2. 下載後，透過 `selenium.NewChromeDriverService()` 中的指定路徑
        - https://godoc.org/github.com/tebeka/selenium#NewChromeDriverService

