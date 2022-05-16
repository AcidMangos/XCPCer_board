package scraper

import (
	"XCPCer_board/model"
	"github.com/gocolly/colly"
	"time"
)

// @Author: Feng
// @Date: 2022/4/7 15:44

//Scrape 爬
func (s *Scraper) Scrape(do func(*colly.Collector) error) ([]KV, error) {
	select {
	case p := <-s.ch:
		// 执行访问
		err := do(p.c)
		if err != nil {
			s.ch <- p
			return nil, err
		}
		// 读取访问结果
		var re []KV
		finish := false
		for !finish {
			select {
			case kv := <-p.ch:
				re = append(re, kv)
			default:
				finish = true
			}
		}
		s.ch <- p
		return re, nil
	case <-time.After(s.timeout):
		return nil, model.ScrapeTimeoutError
	}
}
