package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	beego.Notice("this is notice")
	this.Ctx.WriteString("hello world")
}

func main() {
	beego.Router("/", &MainController{})
	//闭包
	beego.Get("/hello", func(context *context.Context) {
		context.Output.Body([]byte("hee"))
	})
	//无论get post
	beego.Any("/foo", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bar"))
	})
	beego.Run()
}
