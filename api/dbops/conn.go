package dbops

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var (
    dbConn *sql.DB
    err error
)

func init() {
    // 数据库连接
    dbConn, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:33060)/video_server?charset=utf8")
    if err != nil {
        panic(err.Error())
    }
}
