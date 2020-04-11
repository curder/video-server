package dbops

import "log"

// 添加用户
func AddUserCredential(loginName string, password string) error {
    stmtIns, err := dbConn.Prepare("INSERT INTO `users` (`login_name`, `pwd`) VALUES (?, ?)")
    if err != nil {
        return err
    }

    stmtIns.Exec(loginName, password)
    defer stmtIns.Close()
    return nil
}

// 获取用户
func GetUserCredential(loginName string) (string, error) {
    stmtOut, err := dbConn.Prepare("SELECT * FROM `users` WHERE `login_name` = ?")
    if err != nil {
        log.Panicf("%s", err)
        return "", err
    }

    var pwd string
    stmtOut.QueryRow(loginName).Scan(&pwd)
    defer stmtOut.Close()

    return pwd, nil
}
