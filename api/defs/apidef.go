package defs

// 用户结构体
type UserCredential struct {
    UserName string `json:"user_name"` // 用户名
    Password string `json:"password"`  // 密码
}

// 视频结构体
type VideoInfo struct {
    Id string `json:"id"`
    AuthorId int `json:"author_id"`
    Name string `json:"name"`
    DisplayCtime string `json:"display_ctime"`
}