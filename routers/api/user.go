package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rawPracticeNick/models"
	"rawPracticeNick/pkg/app"
	"rawPracticeNick/pkg/e"
	"rawPracticeNick/pkg/util"
)

func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}
	openId := c.Query("openId")
	user, err := models.SelectUserByOpenId(openId)
	if err != nil {
		log.Println("获取用户错误 :" + err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, user)
}

func GetOpenId(c *gin.Context) {
	appG := app.Gin{C: c}
	jsCode := c.Query("jsCode")
	log.Println("jsCode is :", jsCode)
	body, _ := util.GetOpenId(jsCode)
	log.Println("body is :", body)
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
