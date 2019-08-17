package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRouters()
}
func (a *App) initializeRouters() {
	//findUserByIdHandler := a.Router.HandlerFunc(HandleFindUserById)
	//a.Router.HandleFunc("/users/findUserById", loggingHandler(findUserByIdHandler))
	a.Router.Use(loggingHandler)
	a.Router.HandleFunc("/users/findByName", HandleFindUserByName)
	a.Router.HandleFunc("/users/updateUser", HandleUpdateUser)
	a.Router.HandleFunc("/users/addUser", HandleAddUser)
	a.Router.HandleFunc("/users/delUser", HandleDelUser)
}
func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Printf("[%s] %q", r.Method, r.URL.String())
	}
	return http.HandlerFunc(fn)
}
