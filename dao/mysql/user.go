package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"webapp/models"
	"webapp/pkg/snowflake"

	"go.uber.org/zap"
)

// 存放数据库相关的操作，增删改查

var secret = "夏天夏天悄悄过去"

// encryptPassword 加密
func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))                // 先加盐
	return hex.EncodeToString(h.Sum(data)) // 再md5并转为十六进制字符串
}

// Register 将指定用户注册
func Register(user *models.User) (err error) {
	// 1. 先判断当前用户名是否已经被注册过
	if CheckUserIsExist(user.UserName) {
		return ErrorUserExit
	}
	// 2. 明文的密码要加盐加密处理处理才能入库
	password := encryptPassword([]byte(user.Password))
	// user_id
	userID, err := snowflake.GenID()
	if err != nil {
		zap.L().Error("snowflake.GenID failed", zap.Error(err))
		return err
	}
	// 3. 入库
	sqlStr := `insert into user (user_id, username, password) values (?, ?,?)`
	_, err = db.Exec(sqlStr, userID, user.UserName, password)
	return
	// 4. 返回
}

func CheckUserIsExist(username string) bool {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int64
	err := db.Get(&count, sqlStr, username)
	if err != nil && err != sql.ErrNoRows {
		zap.L().Error("query user exist failed", zap.Error(err))
		return true
	}
	return count > 0
}
