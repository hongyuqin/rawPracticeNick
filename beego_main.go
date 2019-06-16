package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}
func (this *MainController) Post() {
	this.Ctx.WriteString("hello world in post")
}

type RegexExpController struct {
	beego.Controller
}

func (this *RegexExpController) Get() {
	this.Ctx.WriteString("In regex mode")
	id := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString(fmt.Sprintf(" id is :%s", id))
}
func main() {
	beego.Router("/", &MainController{})
	//闭包
	/*beego.Get("/hello", func(context *context.Context) {
		context.Output.Body([]byte("hee"))
	})
	//无论get post
	beego.Any("/foo", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bar"))
	})
	beego.SetLogger("file", `{"filename":"test.log"}`)
	//输出文件名和行号
	beego.SetLogFuncCall(true)*/
	beego.Router("/regex/?:id", &RegexExpController{})
	beego.Router("/regex1/:id([0-9]+)", &RegexExpController{})
	beego.Run("127.0.0.1:8081")
}
