package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IncludeAdminRoutes 加载admin的路由
func IncludeAdminRoutes(r *gin.Engine) {
	r.GET("/xxx", func(c *gin.Context) {
		c.String(http.StatusOK, "admin/xxx")
	})
}
