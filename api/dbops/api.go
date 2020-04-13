package dbops

import (
    "database/sql"
    "github.com/curder/video-server/api/defs"
    "github.com/curder/video-server/api/utils"
    "log"
    "time"
)

// 添加用户
func AddUserCredential(loginName string, password string) error {
    stmtIns, err := dbConn.Prepare("INSERT INTO `users` (`login_name`, `pwd`) VALUES (?, ?)")
    if err != nil {
        return err
    }
    defer stmtIns.Close()

    _, err = stmtIns.Exec(loginName, password)
    if err != nil {
        return err
    }

    return nil
}

// 获取用户
func GetUserCredential(loginName string) (string, error) {
    stmtOut, err := dbConn.Prepare("SELECT `pwd` FROM `users` WHERE `login_name` = ?")
    if err != nil {
        log.Panicf("%s", err)
        return "", err
    }

    defer stmtOut.Close()

    var pwd string
    err = stmtOut.QueryRow(loginName).Scan(&pwd)
    if err != nil && err != sql.ErrNoRows {
        return "", err
    }

    return pwd, nil
}

// 删除用户
func DeleteUser(loginName string, pwd string) error {
    stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
    if err != nil {
        log.Printf("DeleteUser error: %s", err)
        return err
    }
    defer stmtDel.Close()

    _, err = stmtDel.Exec(loginName, pwd)
    if err != nil {
        return err
    }

    return nil
}

// 添加视频信息
func AddNewVideo(authorId int, name string) (videoInfo *defs.VideoInfo, error error) {
    // 创建数据主键
    uuid, err := utils.NewUUID()
    if err != nil {
        return nil, err
    }

    ctime := time.Now().Format("Jan 02 2006, 15:04:05")

    stmtIns, err := dbConn.Prepare("INSERT INTO `video_info` (`id`, `author_id`, `name`, `display_ctime`) VALUES(?, ?, ?, ?)")
    if err != nil {
        return nil, err
    }

    defer stmtIns.Close()

    _, err = stmtIns.Exec(uuid, authorId, name, ctime)
    if err != nil {
        return nil, err
    }

    res := &defs.VideoInfo{Id: uuid, Name: name, DisplayCtime: ctime}

    return res, nil
}

//
func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
    // create uuid
    stmtOut, err := dbConn.Prepare("SELECT  author_id, name, display_ctime FROM video_info WHERE id=?")

    var aid int
    var dct string
    var name string

    err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
    if err != nil && err != sql.ErrNoRows {
        return nil, err
    }

    if err == sql.ErrNoRows {
        return nil, nil
    }

    defer stmtOut.Close()

    res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}

    return res, nil
}

func DeleteVideoInfo(vid string) error {
    stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
    if err != nil {
        return err
    }

    _, err = stmtDel.Exec(vid)
    if err != nil {
        return err
    }

    defer stmtDel.Close()
    return nil
}