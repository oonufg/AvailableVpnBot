package domain

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func GetDownloadUrlsFromSource() ([]string, error) {
	downloadUrls := make([]string, 0)
	cScrapper := colly.NewCollector()
	cScrapper.OnHTML(".list", func(e *colly.HTMLElement) {
		url := e.ChildAttr("a", "href")
		if url != "" {
			downloadUrls = append(downloadUrls, url)
		}
	})

	cScrapper.Visit("https://ipspeed.info/freevpn_openvpn.php")
	return downloadUrls, nil
}

func GetCountryByIp(ip string) string {
	cScrapper := colly.NewCollector()
	var country string

	cScrapper.OnHTML(".entry", func(e *colly.HTMLElement) {
		country = strings.Split(e.ChildTexts("p")[1], " : ")[1]
	})

	cScrapper.Visit(fmt.Sprintf("https://geoip.co/?ip=%s", ip))
	return country
}
