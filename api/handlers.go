package main

import (
    "github.com/julienschmidt/httprouter"
    "io"
    "net/http"
)

// 创建用户
func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    //
}

// 用户登录
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    name := p.ByName("user_name")

    io.WriteString(w, name)
}
