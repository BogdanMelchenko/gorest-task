package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	routes "github.com/BogdanMelchenko/gorest-task/application/routes"
	stores "github.com/BogdanMelchenko/gorest-task/application/stores"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var envr *routes.Env

func init() {
	connectionString :=
		fmt.Sprintf("user=%s sslmode=%s host=%s port=%s password=%s dbname=%s",
			os.Getenv("TEST_DB_USERNAME"), os.Getenv("TEST_DB_SSL"), os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_PASSWORD"), os.Getenv("TEST_DB_NAME"))

	psDb, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	muxRouter := mux.NewRouter()
	psds := stores.PostgresDbStore{Db: psDb}

	envr = &routes.Env{Router: muxRouter, TaskStore: &psds, UserStore: &psds}
	routes.InitilizeRoutes(envr)
}

func main() {
	Run(os.Getenv("TEST_APP_PORT"))
}

func Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, envr.Router))
}
