package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"strconv"
	"time"
)

func scrapingCrypto(page *agouti.Page) {

	time.Sleep(2 * time.Second)
	page.FindByXPath("//*[@id=\"tabContainer\"]/div[2]/div[2]/div/div[1]/div[2]/div").Click()
	cryptData := make(map[string]string)

	for i := 0; i < 5; i++ {

		if i >= 1 {
			nextButton := fmt.Sprintf("//*[@id=\"tabContainer\"]/div[2]/div[3]/div/button[%s]", strconv.Itoa(i+2))
			page.FindByXPath(nextButton).Click()
			time.Sleep(2 * time.Second)
		}
		for dataCnt := 1; dataCnt < 20; dataCnt++ {
			currencyName := fmt.Sprintf("//*[@id=\"tabContainer\"]/div[2]/div[2]/div/div[2]/div[%s]/div/div[1]/div[2]/div", strconv.Itoa(dataCnt))
			currencyPrice := fmt.Sprintf("//*[@id=\"tabContainer\"]/div[2]/div[2]/div/div[2]/div[%s]/div/div[2]/div", strconv.Itoa(dataCnt))
			name, _ := page.FindByXPath(currencyName).Text()
			price, _ := page.FindByXPath(currencyPrice).Text()
			cryptData[name] = price
		}
	}
	fmt.Println(cryptData)
}

func openChromedriver() {
	chromeArgs := agouti.ChromeOptions(
		"args", []string{
			"--headless",
			"--disable-gpu",
		})
	chromeExcludeSwitches := agouti.ChromeOptions(
		"excludeSwitches", []string{
			"enable-logging",
		})
	driver := agouti.ChromeDriver(chromeArgs, chromeExcludeSwitches, agouti.Browser("chrome"))

	if err := driver.Start(); err != nil {
		fmt.Printf("Failed to start driver. %s\n", err)
		return
	}
	page, err := driver.NewPage()
	if err != nil {
		fmt.Printf("Failed to open a new page. %s\n", err)
		return
	}
	if err := page.Navigate("https://www.binance.com/ja/markets/coinInfo-"); err != nil {
		fmt.Printf("Failed to navigate to https://readouble.com/laravel/8.x/ja/eloquent.html. %s\n", err)
		return
	}
	scrapingCrypto(page)
}

func main() {
	openChromedriver()
}
