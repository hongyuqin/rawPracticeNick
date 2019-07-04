package main

import (
	"fmt"
	"rawPraticeNick/models"
	"rawPraticeNick/pkg/setting"
	"rawPraticeNick/routers"
)

func init() {
	setting.SetUp()
	models.Setup()
}
func main() {
	r := routers.InitRouter()

	r.Run(fmt.Sprintf(":%d", setting.ServerSetting.HttpPort))
}
