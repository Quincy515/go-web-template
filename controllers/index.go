package controllers

import (
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	zap.L().Error("this is a error")
	zap.L().Debug("this is index handler")
	zap.L().Info("this is a test log")
	time.Sleep(time.Second * 5)
	c.String(http.StatusOK, viper.GetString("app.ver"))
}
