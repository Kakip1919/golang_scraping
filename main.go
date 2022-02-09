package main

import (
	"fmt"
	"github.com/sclevine/agouti"
)

func openChromedriver() {

	driver := agouti.ChromeDriver(agouti.Browser("chrome"))

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

	page.FindByXPath("//*[@id=\"tabContainer\"]/div[2]/div[2]/div/div[1]/div[2]/div").Click()
}

func main() {
	openChromedriver()
}
