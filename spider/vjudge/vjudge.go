package vjudge

import (
	"XCPCer_board/scraper"
)

//ScrapeAll 获得所有结果
func ScrapeAll(uid string) (map[string]int, error) {
	// 请求所有并合并所有
	res, err := scraper.MergeAllResults[string, int](
		GetIntMsg(uid),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}