package initDB

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "log"

// 定义全局变量
var Db *sql.DB

// 初始化链接
func init()  {
	// TODO 注意变量作用域
	var err error
	Db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gin")

	if err != nil {
		log.Panicf("连接数据库失败,嘤嘤嘤\n")
	}
	Db.SetMaxIdleConns(6)
	//最大连接数
	Db.SetMaxOpenConns(6)

}
