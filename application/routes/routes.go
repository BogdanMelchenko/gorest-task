package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type Env struct {
	Router *mux.Router
	Db     *sql.DB
}

func InitilizeRoutes(env *Env) {
	initializeTaskRoutes(env.Router, env.Db)
	initializeUserRoutes(env.Router, env.Db)
}
