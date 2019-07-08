package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"rawPracticeNick/models"
	"rawPracticeNick/pkg/app"
	"rawPracticeNick/pkg/e"
	"rawPracticeNick/pkg/util"
	"rawPracticeNick/service/user_service"
)

func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}
	openId := c.Query("openId")
	user, err := models.SelectUserByOpenId(openId)
	if err != nil {
		logrus.Info("获取用户错误 :" + err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, user)
}

func GetOpenId(c *gin.Context) {
	appG := app.Gin{C: c}
	jsCode := c.Query("jsCode")
	logrus.Info("jsCode is :", jsCode)
	body, _ := util.GetOpenId(jsCode)
	logrus.Info("body is :", body)
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func HomePage(c *gin.Context) {
	appG := app.Gin{C: c}
	homeDetail, err := user_service.Home("xx")
	if err != nil {
		logrus.Error("Home error :", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, homeDetail)
}
