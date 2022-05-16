package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"sync"
	"time"
)

// @Author: Feng
// @Date: 2022/4/8 17:38

//Scraper colly封装
type Scraper struct {
	cb      func(collector *colly.Collector, res *Processor) // 初始化每一个colly controller的回调函数
	ch      chan *Processor                                  // processor复用管道
	threads uint32                                           // 启动的持久化协程数量和processor数量
	timeout time.Duration                                    // scrape超时时间
}

//参数丰富接口
type scraperFunc func(*Scraper)

//WithCallback 带上处理回调函数
func WithCallback(cb func(collector *colly.Collector, res *Processor)) scraperFunc {
	return func(s *Scraper) {
		s.cb = cb
	}
}

//WithThreads 带上并发处理的协程数
func WithThreads(threads uint32) scraperFunc {
	return func(s *Scraper) {
		s.threads = threads
	}
}

//WithTimeout 带上并发处理的超时时间
func WithTimeout(timeout time.Duration) scraperFunc {
	if timeout < time.Second {
		timeout = time.Second
	}
	return func(s *Scraper) {
		s.timeout = timeout
	}
}

//defaultCallback 默认的处理回调函数
func defaultCallback() func(*colly.Collector, *Processor) {
	return func(c *colly.Collector, res *Processor) {
		c.OnRequest(func(r *colly.Request) {
			fmt.Println(r.URL)
			res.Set("Default Callback 1", struct{}{})
		})
	}
}

//NewScraper 构造Scraper
func NewScraper(opts ...scraperFunc) *Scraper {
	// 默认参数
	s := Scraper{
		timeout: 5 * time.Second,
		threads: 5,
		cb:      defaultCallback(),
	}
	// 应用外来参数
	for _, f := range opts {
		if f != nil {
			f(&s)
		}
	}
	s.ch = make(chan *Processor, s.threads)
	// 初始化
	s.init()
	return &s
}

//init 初始化
func (s *Scraper) init() {
	wg := sync.WaitGroup{}
	// 初始化各种On
	for i := uint32(0); i < s.threads; i++ {
		// 初始化并开始监听
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := colly.NewCollector(
				colly.Async(false),
				colly.AllowURLRevisit(),
			)
			p := NewProcessor(c, s.cb)
			s.ch <- p
		}()
		// 启动持久化处理协程
		go newFlusher()
	}
	wg.Wait()
}
