package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
	"net/http"
	"rawPracticeNick/pkg/app"
	"rawPracticeNick/pkg/e"
	"rawPracticeNick/service/topic_service"
	"strconv"
)

func NextTopic(c *gin.Context) {
	appG := app.Gin{C: c}

	topic, err := topic_service.NextTopic("xx")
	if err != nil {
		logrus.Error("NextTopic error :", err)
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

type AnswerReq struct {
	OpenId   string `schema:"open_id"`
	MyAnswer string `schema:"my_answer"`
	TopicId  int    `schema:"topic_id"`
}

func Answer(c *gin.Context) {
	appG := app.Gin{C: c}
	var decoder = schema.NewDecoder()
	req := &AnswerReq{}
	if err := decoder.Decode(&req, c.Request.URL.Query()); err != nil {
		logrus.Error("decode error :", err)
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	resp, err := topic_service.Answer(req.OpenId, req.TopicId, req.MyAnswer)
	if err != nil {
		logrus.Error("answer error :", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, resp)
}
