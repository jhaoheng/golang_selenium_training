package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const website = "http://spidertestsite/"

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	proxyHttp()
}

func handleHeader(resp *http.Response) {
	fmt.Println(resp.Header)
}

func handleBody(resp *http.Response) {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func handleCookies(resp *http.Response) {
	cookies := resp.Cookies()
	fmt.Println(cookies)
}

// proxy
func proxyHttp() {
	//creating the proxyURL
	proxyStr := "http://proxyserver:8080"
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		log.Println(err)
	}

	// creating the URL to be loaded through the proxy
	urlStr := website
	url, err := url.Parse(urlStr)
	if err != nil {
		log.Println(err)
	}

	//adding the proxy settings to the Transport object
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	//adding the Transport object to the http Client
	client := &http.Client{
		Transport: transport,
	}

	//generating the HTTP GET request
	request, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Println(err)
	}

	//calling the URL
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}

	//getting the response
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	//printing the response
	log.Println(string(data))
}
