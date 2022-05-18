package codeforces

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"strings"
)

var (
	intScraper = scraper.NewScraper(
		scraper.WithCallback(intCallback),
		scraper.WithThreads[int](2),
	)
)

const (
	// key
	// 个人总过题数
	problemPassAmountKey = "codeForces_problem_pass_amount"
	// 个人最后一月过题数
	lastMonthPassAmount = "codeForces_last_month_problem_pass_amount"

	// CF finder关键词
	// 个人总过题数
	problemPassKeyWord = "all"
	//个人最后一月过题数
	lastMonthPassKeyWord = "month"
)

//problemPassAmountHandler 获取cf个人总过题数
func problemPassAmountHandler(doc *goquery.Selection) string {
	retStr := doc.Find(fmt.Sprintf("div[style=\"position: relative;\"] #pageContent ._UserActivityFrame_frame"+
		" .roundbox.userActivityRoundBox ._UserActivityFrame_footer"+
		" ._UserActivityFrame_countersRow ._UserActivityFrame_counter:contains(solved):contains(%v)"+
		" ._UserActivityFrame_counterValue", problemPassKeyWord)).First().Text()
	return strings.Split(retStr, " ")[0]
	//"1000 problems" -> "1000"
}

//lastMonthAmountHandler 获取cf个人上个月总过题数
func lastMonthAmountHandler(doc *goquery.Selection) string {
	retStr := doc.Find(fmt.Sprintf("div[style=\"position: relative;\"] #pageContent ._UserActivityFrame_frame"+
		" .roundbox.userActivityRoundBox ._UserActivityFrame_footer"+
		" ._UserActivityFrame_countersRow ._UserActivityFrame_counter:contains(solved):contains(%v)"+
		" ._UserActivityFrame_counterValue", lastMonthPassKeyWord)).First().Text()
	return strings.Split(retStr, " ")[0]
}

func intCallback(c *colly.Collector, res *scraper.Processor) {
	c.OnHTML("#body", func(e *colly.HTMLElement) {
		//fmt.Println(r.DOM.First().Text())
		res.Set(problemPassAmountKey, strToInt(e.DOM, problemPassAmountHandler))
		res.Set(lastMonthPassAmount, strToInt(e.DOM, lastMonthAmountHandler))
	})
}

//GetIntMsg 对外暴露函数，获取int信息
func GetIntMsg(uid string) ([]scraper.KV, error) {
	return intScraper.Scrape(getPersonPage(uid))
}
