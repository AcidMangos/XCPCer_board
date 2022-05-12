package atcoder

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

//---------------------------------------------------------------------//
// atCoder个人信息 //
//---------------------------------------------------------------------//
//  Key

const (
	//key

	RatingKey     = "atc_rating"
	contestSumKey = "atc_contest_sum"
	rankKey       = "atc_rank"

	//keyword
)

var (
	mainScraper = scraper.NewScraper[int](
		scraper.WithCallback(mainCallback),
		scraper.WithThreads[int](2),
	)
)

//mainCallback 处理个人主页的回调函数
func mainCallback(c *colly.Collector, res *scraper.Processor[int]) {
	//用goquery
	c.OnHTML("body #main-div #main-container .row .col-md-9.col-sm-12 .dl-table.mt-2",
		func(element *colly.HTMLElement) {
			ret := element.DOM.Find(fmt.Sprintf("tr:nth-child(2) td span:first-child")).First().Text()
			if num, err := strconv.Atoi(ret); err == nil {
				res.Set(RatingKey, num)
			}
			res.Set(contestSumKey, contestSumHandler(element.DOM))
			res.Set(rankKey, rankHandler(element.DOM))
		},
	)
}

//getAtCoderBaseUrl 获取个人主页URL
func getAtCoderBaseUrl(atCoderId string) string {
	return "https://atcoder.jp/users/" + atCoderId
}

//ratingHandler 获取个人rating
func ratingHandler(doc *goquery.Selection) int {

	return -1
}

//rankHandler 获取个人rating排名
func rankHandler(doc *goquery.Selection) int {
	ret := doc.Find(fmt.Sprintf("tr:nth-child(1) td")).First().Text()
	ret = strings.Split(ret, "th")[0]
	if num, err := strconv.Atoi(ret); err == nil {
		return num
	}
	return -1
}

//contestSumHandler 获取比赛场数
func contestSumHandler(doc *goquery.Selection) int {
	ret := doc.Find(fmt.Sprintf("tr:nth-child(4) td")).First().Text()
	if num, err := strconv.Atoi(ret); err == nil {
		return num
	}
	return -1
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//FetchMainPage 抓取个人主页页面所有
func FetchMainPage(uid string) scraper.Results[int] {
	return mainScraper.Scrape(getAtCoderBaseUrl(uid))
}
