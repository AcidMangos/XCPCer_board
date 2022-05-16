package nowcoder

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"strconv"
)

// @Author: Feng
// @Date: 2022/4/11 16:17

//-------------------------------------------------------------------------------------------//
// 基础方法
//-------------------------------------------------------------------------------------------//

// 牛客finder存储Key
const (
	// 个人主页
	ratingKey        = "nowcoder_rating"
	rankingKey       = "nowcoder_ranking"
	contestAmountKey = "nowcoder_attend_contest_amount"

	// 个人主页selector关键字
	ratingKeyWord        = "Rating"
	ratingRankingKeyWord = "Rating排名"
	contestAmountKeyWord = "次比赛"
)

var (
	mainScraper = scraper.NewScraper(
		scraper.WithCallback(mainCallback),
		scraper.WithThreads(2),
	)
)

//mainCallback 处理牛客个人主页的回调函数
func mainCallback(c *colly.Collector, res *scraper.Processor) {
	//用goquery
	c.OnHTML(".nk-container.acm-container .nk-container .nk-main.with-profile-menu.clearfix .my-state-main",
		func(element *colly.HTMLElement) {
			uid := element.Request.Ctx.Get("uid")
			// rating
			num, err := strconv.Atoi(element.DOM.Find(fmt.Sprintf(".my-state-item:contains(%v) .state-num.rate-score5",
				ratingKeyWord)).First().Text())
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			}
			res.Set(getRatingKey(uid), num)
			// 排名
			num, err = strconv.Atoi(element.DOM.Find(getNowCoderContestBaseFindRule(ratingRankingKeyWord)).First().Text())
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			}
			res.Set(getRankingKey(uid), num)
			// 过题数
			num, err = strconv.Atoi(element.DOM.Find(getNowCoderContestBaseFindRule(contestAmountKeyWord)).First().Text())
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			}
			res.Set(getContestAmountKey(uid), num)
		},
	)

}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//fetchMainPage 抓取个人主页页面所有
func fetchMainPage(uid string) ([]scraper.KV, error) {
	return mainScraper.Scrape(func(c *colly.Collector) error {
		ctx := colly.NewContext()
		ctx.Put("uid", uid)
		err := c.Request("GET", getContestProfileUrl(uid), nil, ctx, nil)
		if err != nil {
			log.Errorf("scraper error %v", err)
			return err
		}
		return nil
	})
}
