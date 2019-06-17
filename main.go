package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 定义一个 Person 结构体，用来绑定 url query
type Person struct {
	Name     string    `form:"name"` // 使用成员变量标签定义对应的参数名
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func startPage(c *gin.Context) {
	var person Person
	// 将 url 查询参数和person绑定在一起
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}
func startPage1(c *gin.Context) {
	var person Person
	// 绑定到 person
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	c.String(200, "Success")
}
func startPage2(c *gin.Context) {
	var person Person
	// 绑定到 person
	if err := c.ShouldBindJSON(&person); err == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	} else {
		log.Println("error : ", err)
	}

	c.String(200, "Success")
}
func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		// 获取post过来的message内容
		// 获取的所有参数内容的类型都是 string
		message := c.PostForm("message")
		// 如果不存在，使用第二个当做默认内容
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	//数据绑定 get
	router.Any("/testing", startPage)
	//绑定表单
	router.POST("/test1", startPage1)
	//绑定json
	router.POST("/testjson", startPage2)
	router.Run(":8080")
}
