package atcoder

import (
	"XCPCer_board/scraper"
	"fmt"
	log "github.com/sirupsen/logrus"
)

var (
	// 爬取函数
	fetchers = []func(uid string) ([]scraper.KV, error){
		fetchMainPage,
		fetchConPage,
	}
)

//scrape 拉取atCoder的所有结果
func scrape(uid string) (res []scraper.KV) {
	// 请求所有
	for _, f := range fetchers {
		// 请求
		kvs, err := f(uid)
		if err != nil {
			log.Errorf("atcoder Fetcher Error %v", err)
			continue
		}
		res = append(res, kvs...)
	}
	fmt.Println(len(res))
	fmt.Println(res)

	return res
}

//Flush 刷新atCoder某用户信息
func Flush(uid string) {
	// 拉出所有kv对
	kvs := scrape(uid)
	// 向持久化处理协程注册持久化处理函数
	scraper.JustLog(kvs)
}
