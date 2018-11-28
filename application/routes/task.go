package routes

import (
	"database/sql"
	"net/http"

	handlers "github.com/BogdanMelchenko/gorest-task/application/handlers"
	"github.com/gorilla/mux"
)

func initializeTaskRoutes(router *mux.Router, db *sql.DB) {

	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasks(w, r, db)
	}).Methods("GET")

	router.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTask(w, r, db)
	}).Methods("POST")

	router.HandleFunc("/task/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTask(w, r, db)
	}).Methods("GET")

	router.HandleFunc("/task/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateTask(w, r, db)
	}).Methods("PUT")

	router.HandleFunc("/task/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteTask(w, r, db)
	}).Methods("DELETE")
}
