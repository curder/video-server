package dbops

import "testing"

func TestMain(m *testing.M) {
    clearTables()
    m.Run() // 跑所有的test
    clearTables()
}

// 清空表数据
func clearTables() {
    dbConn.Exec("truncate `users`")
    dbConn.Exec("truncate `video_info`")
    dbConn.Exec("truncate `comments`")
    dbConn.Exec("truncate `sessions`")
}

// 按照自定义流程跑test
func TestUserWorkFlow(t *testing.T) {
    t.Run("Add", testAddUser)
    t.Run("Get", testGetUser)
    t.Run("Delete", testDeleteUser)
    t.Run("ReGet", testReGetUser)
}

// 添加用户
func testAddUser(t *testing.T) {
    err := AddUserCredential("curder", "123456")
    if err != nil {
        t.Errorf("Error of addUser: %v", err)
    }
}

// 获取用户
func testGetUser(t *testing.T) {
    credential, err := GetUserCredential("curder")
    t.Log(credential)
    if credential != "123456" || err != nil {
        t.Errorf("Error of getUser: %v", err)
    }

}

// 删除用户
func testDeleteUser(t *testing.T) {
    err := DeleteUser("curder", "123456")
    if err != nil {
        t.Errorf("Error of DeleteUser: %v", err)
    }
}

// 获取用户
func testReGetUser(t *testing.T) {
    pwd, err := GetUserCredential("curder")

    if err != nil {
        t.Errorf("Error of RegetUser: %v", err)
    }

    if pwd != "" {
        t.Errorf("Deleting user test failed")
    }
}
