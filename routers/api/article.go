package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"rawPracticeNick/pkg/app"
	"rawPracticeNick/pkg/e"
	"rawPracticeNick/pkg/gredis"
)

func GetArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	body, err := gredis.Get("name")
	if err != nil {
		log.Error("error ", err)

	} else {
		log.Info("name : ", string(body))
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"success": "success",
	})
}
