package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"rawPracticeNick/common"
	"rawPracticeNick/models"
	"rawPracticeNick/pkg/app"
	"rawPracticeNick/pkg/e"
	"rawPracticeNick/pkg/gredis"
	"rawPracticeNick/pkg/util"
	"rawPracticeNick/service/user_service"
	"strconv"
)

func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}
	openId := c.Query("accessToken")
	user, err := models.SelectUserByOpenId(openId)
	if err != nil {
		logrus.Info("获取用户错误 :" + err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, user)
}

type Resp struct {
	SessionKey string `json:"session_key"`
	OpenId     string `json:"openid"`
}

func GetAccessToken(c *gin.Context) {
	appG := app.Gin{C: c}
	jsCode := c.Query("code")
	logrus.Info("jsCode is :", jsCode)
	body, _ := util.GetOpenId(jsCode)
	logrus.Info("body is :", body)

	resp := &Resp{}
	if err := json.Unmarshal([]byte(body), resp); err != nil {
		logrus.Error("marshal error :", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	}
	//1.保存用户
	exist, err := gredis.SIsmember(common.OPENID_SET, resp.OpenId)
	if err != nil {
		logrus.Error("redis error :", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	if !exist {
		//不存在该用户 插入redis,更新数据库
		if err = gredis.SAdd(common.OPENID_SET, resp.OpenId); err != nil {
			logrus.Error("redis error :", err)
			appG.Response(http.StatusInternalServerError, e.ERROR, nil)
			return
		}
		if err = models.AddUser(models.User{
			OpenId:       resp.OpenId,
			Region:       common.REGION_SZ,
			ExamType:     common.EXAM_TYPE_CIVIL,
			DailyNeedNum: 100,
		}); err != nil {
			logrus.Error("addUser error :", err)
			appG.Response(http.StatusInternalServerError, e.ERROR, nil)
			return
		}
	}
	//2.返回openId
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"accessToken": resp.OpenId,
	})
	return
}

func HomePage(c *gin.Context) {
	appG := app.Gin{C: c}
	openId := c.Query("accessToken")
	homeDetail, err := user_service.Home(openId)
	if err != nil {
		logrus.Error("Home error :", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, homeDetail)
}

func Plan(c *gin.Context) {
	appG := app.Gin{C: c}
	openId := c.Query("accessToken")
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
	openId := c.Query("accessToken")
	user, err := models.SelectUserByOpenId(openId)
	if err != nil || user == nil {
		logrus.Error("findUserByOpenId openId error :", openId, err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	data := make(map[string]interface{})
	data["region"] = user.Region
	data["exam_type"] = user.ExamType
	data["daily_need_num"] = user.DailyNeedNum
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
