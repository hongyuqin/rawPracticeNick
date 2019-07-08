package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"rawPracticeNick/pkg/app"
	"rawPracticeNick/pkg/e"
	"rawPracticeNick/service/topic_service"
	"strconv"
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

func Collect(c *gin.Context) {
	appG := app.Gin{C: c}
	openId := c.Query("openId")
	topicIdStr := c.Query("topic_id")
	topicId, err := strconv.Atoi(topicIdStr)
	if err != nil {
		logrus.Error("no topic_id")
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	if err = topic_service.Collect(openId, topicId); err != nil {
		logrus.Error("collect error :", openId, topicId)
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
