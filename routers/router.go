package routers

import (
	"github.com/gin-gonic/gin"
	"rawPraticeNick/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.GET("/article", api.GetArticle)
	r.POST("/upload", api.UploadFile)
	r.GET("/getUser", api.GetUser)
	return r
}
