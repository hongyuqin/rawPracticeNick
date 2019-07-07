package routers

import (
	"github.com/gin-gonic/gin"
	"rawPracticeNick/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//测试接口
	r.GET("/article", api.GetArticle)
	r.POST("/upload", api.UploadFile)
	r.GET("/getUser", api.GetUser)
	r.GET("/getOpenId", api.GetOpenId)
	//页面接口
	r.GET("/home", api.HomePage)
	return r
}
