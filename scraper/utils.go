package scraper

import "github.com/gocolly/colly"

// @Author: Feng
// @Date: 2022/5/16 21:29

//Parse 解析context为KV列表切片
func Parse(ctx *colly.Context, ignore map[string]struct{}) (re []KV) {
	ret := ctx.ForEach(func(k string, v interface{}) interface{} {
		return KV{
			Key: k,
			Val: v,
		}
	})
	for _, r := range ret {
		kv, ok := r.(KV)
		if !ok {
			continue
		}
		if _, ok := ignore[kv.Key]; !ok {
			re = append(re, kv)
		}
	}
	return re
}
