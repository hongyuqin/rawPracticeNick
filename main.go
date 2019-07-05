package main

import (
	"fmt"
	"rawPracticeNick/models"
	"rawPracticeNick/pkg/setting"
	"rawPracticeNick/routers"
)

func init() {
	setting.SetUp()
	models.Setup()
}
func main() {
	r := routers.InitRouter()

	r.Run(fmt.Sprintf(":%d", setting.ServerSetting.HttpPort))
}