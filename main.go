package main

import (
	"XCPCer_board/config"
	"XCPCer_board/db/mysql"
	"XCPCer_board/model"
	"XCPCer_board/spider/luogu"
	"fmt"
	_ "github.com/FengZhg/go_tools/gin_logrus"
)

// 主入口函数
func main() {

	//fmt.Println(luogu.ScrapeAll(model.TestLuoGuIdLYF))
	config.InitAll()
	//fmt.Println(config.Config)
	mysql.InitDB()
	fmt.Println(luogu.InsertSql(model.TestLuoGuIdLYF))

}
