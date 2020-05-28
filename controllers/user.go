package controllers

import (
	"webapp/dao/mysql"
	"webapp/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	// 1. 提取用户提交的注册信息
	// 2. 并进行数据校验
	var sd models.SignUpForm
	if err := c.ShouldBindJSON(&sd); err != nil {
		// 返回参数错误的响应
		zap.L().Error("invalid params", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	// 3. 保存到数据库
	u := &models.User{
		UserName: sd.UserName,
		Password: sd.Password,
	}
	if err := mysql.Register(u); err != nil {
		zap.L().Error("sign up user failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 4. 返回响应
	ResponseSuccess(c, nil)
}
