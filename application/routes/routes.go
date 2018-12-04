package routes

import (
	stores "github.com/BogdanMelchenko/gorest-task/application/stores"
	"github.com/gorilla/mux"
)

type Env struct {
	Router    *mux.Router
	UserStore stores.UserStore
	TaskStore stores.TaskStore
}

func InitilizeRoutes(env *Env) {
	initializeTaskRoutes(*env)
	initializeUserRoutes(*env)
}
