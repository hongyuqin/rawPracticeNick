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
	"strconv"
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

func Plan(c *gin.Context) {
	appG := app.Gin{C: c}
	openId := c.Query("openId")
	region := c.Query("region")
	examType := c.Query("exam_type")
	dailyNeedNumStr := c.Query("daily_need_num")
	dailyNeedNum, err := strconv.Atoi(dailyNeedNumStr)
	if err != nil {
		logrus.Error("dailyNeedNum trans error :", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	err = user_service.Plan(openId, region, examType, dailyNeedNum)
	if err != nil {
		logrus.Error("Plan error :", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetPlan(c *gin.Context) {
	appG := app.Gin{C: c}
	openId := c.Query("openId")
	user, err := models.SelectUserByOpenId(openId)
	if err != nil || user == nil {
		logrus.Error("findUserByOpenId error :", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	data := make(map[string]interface{})
	data["region"] = user.Region
	data["exam_type"] = user.ExamType
	data["daily_need_num"] = user.DailyNeedNum
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
