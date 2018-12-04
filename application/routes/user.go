package routes

import (
	"net/http"

	handlers "github.com/BogdanMelchenko/gorest-task/application/handlers"
)

func initializeUserRoutes(env Env) {
	env.Router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUsers(w, r, env.UserStore)
	}).Methods("GET")

	env.Router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUser(w, r, env.UserStore)
	}).Methods("POST")

	env.Router.HandleFunc("/user/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUser(w, r, env.UserStore)
	}).Methods("GET")

	env.Router.HandleFunc("/user/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateUser(w, r, env.UserStore)
	}).Methods("PUT")

	env.Router.HandleFunc("/user/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteUser(w, r, env.UserStore)
	}).Methods("DELETE")
}
