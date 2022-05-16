package scraper

import (
	"XCPCer_board/dao"
	"context"
	log "github.com/sirupsen/logrus"
)

// @Author: Feng
// @Date: 2022/5/12 13:36

var (
	// 持久化任务分配管道
	flushCh = make(chan interface{})
)

type redisRequest struct {
	kvs []KV
}

type dbRequest struct {
	query string
	args  []interface{}
}

//newFlushProcessor 新持久化处理器
func newFlushProcessor() {
	for {
		i := <-flushCh
		switch v := i.(type) {
		case *redisRequest:
			internalFlushRedis(v)
		case *dbRequest:
			internalFlushDB(v)
		}
	}
}

//internalFlushRedis 内部刷新redis数据
func internalFlushRedis(req *redisRequest) {
	for _, kv := range req.kvs {
		// 底层库实现了自动重试
		err := dao.RedisClient.Set(context.Background(), kv.Key, kv.Val, 0).Err()
		if err != nil {
			log.Errorf("internal flush redis error %v", err)
		}
	}

}

//internalFlushDB 内部刷新db内数据
func internalFlushDB(req *dbRequest) {
	_, err := dao.DBClient.Exec(req.query, req.args...)
	if err != nil {
		log.Errorf("internal flush db error %v", err)
	}
}

//FlushRedis 刷新Redis
func FlushRedis(kvs []KV) {
	flushCh <- &redisRequest{
		kvs: kvs,
	}
}

//FlushDB 刷新DB
func FlushDB(query string, args ...interface{}) {
	flushCh <- &dbRequest{
		query: query,
		args:  args,
	}
}
