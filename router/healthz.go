package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sixinshuier/go-web/pkg/app"
	code "github.com/sixinshuier/go-web/pkg/code"
	"net/http"
	"os"
)

// Healthz godoc
// @Tags healthz
// @Summary  Get config
// @Description 健康检查
// @Accept  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /healthz [get]
func Healthz(c *gin.Context) {
	appG := app.Gin{C: c}
	appEnv := os.Getenv("VERSION")
	if appEnv == "" {
		appEnv = "1.1"
	}
	c.Request.Header.Set("VERSION", appEnv)
	appG.Response(http.StatusOK, code.SUCCESS, map[string]interface{}{
		"healthz": "HEALTHZ",
	}, c.Request.Header)
}
