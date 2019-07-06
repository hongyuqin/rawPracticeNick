package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"rawPracticeNick/models"
	"rawPracticeNick/pkg/gredis"
	"rawPracticeNick/pkg/setting"
	"rawPracticeNick/routers"
)

func init() {
	setting.SetUp()
	models.Setup()
	if err := gredis.Setup(); err != nil {
		logrus.Panic("redis init error")
	}
}
func main() {
	r := routers.InitRouter()

	r.Run(fmt.Sprintf(":%d", setting.ServerSetting.HttpPort))
}
