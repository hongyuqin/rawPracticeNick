package api

import (
	"../../pkg/app"
	"../../pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
