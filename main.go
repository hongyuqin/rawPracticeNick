package main

import (
	"./models"
	"./pkg/setting"
	"./routers"
	"fmt"
	"net/http"
)

func init() {
	setting.SetUp()
	models.Setup()
}
func main() {
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	server.ListenAndServe()
}
