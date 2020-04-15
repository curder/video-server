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

// 获取视频信息
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

// 删除视频信息
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

// 添加评论
func AddNewComment(videoId string, authorId int, content string) error {
    id, err := utils.NewUUID()
    if err != nil {
        return err
    }
    stmtIns, err := dbConn.Prepare("INSERT INTO `comments` (`id`, `video_id`, `author_id`, `content`) VALUES (?, ?, ?, ?)")
    if err != nil {
        return err
    }

    _, err = stmtIns.Exec(id, videoId, authorId, content)
    if err != nil {
        return err
    }

    defer stmtIns.Close()
    return nil
}

func ListComments(videoId string, from, to int) ([]*defs.Comment, error) {
    stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)
		ORDER BY comments.time DESC`)

    var res []*defs.Comment

    rows, err := stmtOut.Query(videoId, from, to)
    if err != nil {
        return res, err
    }

    for rows.Next() {
        var id, name, content string
        if err := rows.Scan(&id, &name, &content); err != nil {
            return res, err
        }

        c := &defs.Comment{Id: id, VideoId: videoId, AuthorId: name, Content: content}
        res = append(res, c)
    }

    defer stmtOut.Close()

    return res, nil
}
