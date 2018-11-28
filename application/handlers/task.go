package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	model "github.com/BogdanMelchenko/gorest-task/application/model"
	util "github.com/BogdanMelchenko/gorest-task/application/util"
	"github.com/gorilla/mux"
)

func GetTask(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	p := model.Task{ID: id}
	if err := p.GetTask(DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			util.RespondWithError(w, http.StatusNotFound, "Task not found")
		default:
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	util.RespondWithJSON(w, http.StatusOK, p)
}

func GetTaskOfUser(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	vars := mux.Vars(r)
	ownerID, err := strconv.Atoi(vars["owner_id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid owner ID")
		return
	}
	taskID, err := strconv.Atoi(vars["task_id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid owner ID")
		return
	}

	p := model.Task{ID: taskID, Owner_Id: ownerID}
	if err := p.GetTaskOfUser(DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			util.RespondWithError(w, http.StatusNotFound, "Task not found")
		default:
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	util.RespondWithJSON(w, http.StatusOK, p)
}

func GetTasks(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	v := r.URL.Query()
	titleFilterString := v.Get("titleFilter")
	var tasks []model.Task
	var err error
	if titleFilterString != "" {
		tasks, err = model.GetTasksFilteredByTitle(DB, titleFilterString)
	} else {

		tasks, err = model.GetTasks(DB)
	}
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, tasks)
}

func GetTasksOfUser(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	v := r.URL.Query()

	titleFilter := v.Get("titleFilter")

	vars := mux.Vars(r)
	ownerID, err := strconv.Atoi(vars["owner_id"])

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "bad owner id")
		return
	}

	tasks, err := model.GetTasksOfUser(DB, ownerID, titleFilter)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	var t model.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := t.CreateTask(DB); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, t)
}

func UpdateTask(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	var t model.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	t.ID = id

	if err := t.UpdateTask(DB); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, t)
}

func DeleteTask(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	t := model.Task{ID: id}
	if err := t.DeleteTask(DB); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
