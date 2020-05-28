package routers

import (
	"webapp/controllers"
	"webapp/logger"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

// SetupRouters 配置项目路由信息
func SetupRouters() *gin.Engine {
	gin.SetMode(viper.GetString("app.mode"))
	//r := gin.Default()
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(false))
	r.GET("/index", controllers.IndexHandler)

	r.POST("/signup", controllers.SignUpHandler)

	//IncludeAdminRoutes(r)
	return r
}
