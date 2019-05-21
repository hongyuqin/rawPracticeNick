package main

import "fmt"

//http://c.biancheng.net/view/79.html
//类型中嵌入其他类型或者结构体来实现  来实现  接口的方法
// 一个服务需要满足能够开启和写日志的功能
type Service interface {
	Start()     //开启服务
	Log(string) //日志输出
}

//日志器
type Logger struct {
}

//实现service的Log()方法
func (g *Logger) Log(l string) {
	fmt.Println("log")
}

//游戏服务
type GameService struct {
	Logger
}

//实现Service的start方法
func (f *GameService) Start() {
	fmt.Println("start")
}
func main() {
	g := new(GameService)
	g.Log("dd")
	g.Start()
}
