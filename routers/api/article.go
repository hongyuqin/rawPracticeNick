package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rawPracticeNick/pkg/app"
	"rawPracticeNick/pkg/e"
)

func GetArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"success": "success",
	})
}
