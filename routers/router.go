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
	//开始答题
	r.GET("/nextTopic", api.NextTopic)
	//9.收藏
	r.GET("/collect", api.Collect)
	//7.修改计划
	r.GET("/plan", api.Plan)
	//6.获取计划
	r.GET("/getPlan", api.GetPlan)
	//7.提交答案
	r.GET("/answer", api.Answer)
	return r
}
