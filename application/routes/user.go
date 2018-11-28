package routes

import (
	"database/sql"
	"net/http"

	handlers "github.com/BogdanMelchenko/gorest-task/application/handlers"
	"github.com/gorilla/mux"
)

func initializeUserRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUsers(w, r, db)
	}).Methods("GET")

	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUser(w, r, db)
	}).Methods("POST")

	router.HandleFunc("/user/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUser(w, r, db)
	}).Methods("GET")

	router.HandleFunc("/user/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateUser(w, r, db)
	}).Methods("PUT")

	router.HandleFunc("/user/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteUser(w, r, db)
	}).Methods("DELETE")

	router.HandleFunc("/user/{owner_id:[0-9]+}/tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasksOfUser(w, r, db)
	}).Methods("GET")

	router.HandleFunc("/user/{owner_id:[0-9]+}/task/{task_id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTaskOfUser(w, r, db)
	}).Methods("GET")

}
