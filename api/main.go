package main

import (
    "github.com/julienschmidt/httprouter"
    "log"
    "net/http"
)

// 注册处理方法
func RegisterHandlers() *httprouter.Router {
    router := httprouter.New()

    router.POST("/user", CreateUser) // 创建用户
    router.POST("/user/:user_name", Login) // 用户登录

    return router
}

func main() {
    r := RegisterHandlers()

    log.Fatal(http.ListenAndServe(":8888", r))
}