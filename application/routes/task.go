package routes

import (
	"net/http"

	handlers "github.com/BogdanMelchenko/gorest-task/application/handlers"
)

func initializeTaskRoutes(env Env) {

	env.Router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasks(w, r, env.TaskStore)
	}).Methods("GET")

	// env.Router.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.CreateTask(w, r, env.TaskStore)
	// }).Methods("POST")

	env.Router.HandleFunc("/task/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTask(w, r, env.TaskStore)
	}).Methods("GET")

	env.Router.HandleFunc("/task/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateTask(w, r, env.TaskStore)
	}).Methods("PUT")

	env.Router.HandleFunc("/task/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteTask(w, r, env.TaskStore)
	}).Methods("DELETE")

	env.Router.HandleFunc("/user/{owner_id:[0-9]+}/tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasksOfUser(w, r, env.TaskStore)
	}).Methods("GET")

	env.Router.HandleFunc("/user/{owner_id:[0-9]+}/task/{task_id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTaskOfUser(w, r, env.TaskStore)
	}).Methods("GET")
}
