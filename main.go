package main

import (
	"./models"
	"./routers"
	"net/http"
)

func init() {
	models.Setup()
}
func main() {
	routersInit := routers.InitRouter()
	server := &http.Server{
		Handler: routersInit,
	}
	server.ListenAndServe()
}
