package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rawPraticeNick/models"
	"rawPraticeNick/pkg/app"
	"rawPraticeNick/pkg/e"
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
