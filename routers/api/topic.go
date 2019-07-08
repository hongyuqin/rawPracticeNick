package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"rawPracticeNick/pkg/app"
	"rawPracticeNick/pkg/e"
	"rawPracticeNick/service/topic_service"
)

func BeginAnswer(c *gin.Context) {
	appG := app.Gin{C: c}

	topic, err := topic_service.BeginAnswer("xx", "REGION_COUNTRY", "常识判断", "政治")
	if err != nil {
		logrus.Error("BeginAnswer error :", err)
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, topic)
}
