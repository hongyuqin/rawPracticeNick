package routers

import (
	"./api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.GET("/article", api.GetArticle)
	return r
}
