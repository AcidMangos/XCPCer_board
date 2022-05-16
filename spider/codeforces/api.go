package codeforces

import (
	"XCPCer_board/scraper"
	"encoding/json"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

// @Author: Feng
// @Date: 2022/5/12 22:41

const (
	// 个人rating
	ratingKey = "rating"
	// 个人历史最高rating
	maxRatingKey = "max_rating"
	//当前rating所对应的等级（红名、紫名...)
	rankingNameKey = "ranking_name"
	//最大rating所对应的等级（红名、紫名...)
	maxRankingNameKey = "max_ranking_name"
)

var (
	apiScraper = scraper.NewScraper(
		scraper.WithCallback(userInfoCallback),
	)
)

//userInfoCallback 处理codeforces的api
func userInfoCallback(c *colly.Collector, res *scraper.Processor) {
	c.OnScraped(func(r *colly.Response) {
		rsp := &UserInfo{}
		err := json.Unmarshal(r.Body, rsp)
		if err != nil {
			log.Errorf("Codeforces User Info Unmarshal Error %v", err)
			return
		}
		if rsp.GetStatus() != "OK" || len(rsp.GetInfos()) != 1 {
			log.Errorf("Response: %v Infos Length: %v", rsp.GetStatus(), len(rsp.GetInfos()))
			return
		}
		res.Set(ratingKey)
	})
}
