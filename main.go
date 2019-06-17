package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloWorldGet(c *gin.Context) {
	c.String(http.StatusOK, "hello world in get!")
}
func helloWorldPost(c *gin.Context) {
	c.String(http.StatusOK, "hello world in post!")
}
func fetchId(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, fmt.Sprintf("id is :%s\n", id))
}
func action1(c *gin.Context) {
	c.String(http.StatusOK, "action 1")
}

func main() {
	fmt.Println("hello world")
	//Restful路由
	router := gin.Default()
	//Restful路由 get
	router.GET("/helloworld", helloWorldGet)
	//Restful路由 post
	router.POST("/helloworld", helloWorldPost)
	//不支持正则路由
	//获取path的参数
	router.GET("/param/:id", fetchId)
	//组路由
	group1 := router.Group("g1")
	{
		group1.GET("/action1", action1)
	}
	//服务启动
	router.Run("127.0.0.1:8082")
}
