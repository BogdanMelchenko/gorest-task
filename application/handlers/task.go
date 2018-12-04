package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	model "github.com/BogdanMelchenko/gorest-task/application/model"
	"github.com/BogdanMelchenko/gorest-task/application/stores"
	util "github.com/BogdanMelchenko/gorest-task/application/util"
	"github.com/gorilla/mux"
)

func GetTask(w http.ResponseWriter, r *http.Request, store stores.TaskStore) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	t := model.Task{ID: id}
	if err := store.GetTask(&t); err != nil {
		switch err {
		case sql.ErrNoRows:
			util.RespondWithError(w, http.StatusNotFound, "Task not found")
		default:
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	util.RespondWithoutError(w, http.StatusOK, t, r.Header.Get("Content-type"))
}

func GetTaskOfUser(w http.ResponseWriter, r *http.Request, store stores.TaskStore) {
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

	t := model.Task{ID: taskID, OwnerID: ownerID}
	if err := store.GetTaskOfUser(&t); err != nil {
		switch err {
		case sql.ErrNoRows:
			util.RespondWithError(w, http.StatusNotFound, "Task not found")
		default:
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	util.RespondWithoutError(w, http.StatusOK, t, r.Header.Get("Content-type"))
}

func GetTasks(w http.ResponseWriter, r *http.Request, store stores.TaskStore) {
	v := r.URL.Query()
	titleFilterString := v.Get("titleFilter")
	var tasks []model.Task
	var err error
	if titleFilterString != "" {
		tasks, err = store.GetTasksFilteredByTitle(titleFilterString)
	} else {

		tasks, err = store.GetTasks()
	}
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithoutError(w, http.StatusOK, tasks, r.Header.Get("Content-type"))
}

func GetTasksOfUser(w http.ResponseWriter, r *http.Request, store stores.TaskStore) {
	v := r.URL.Query()

	titleFilter := v.Get("titleFilter")

	vars := mux.Vars(r)
	ownerID, err := strconv.Atoi(vars["owner_id"])

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "bad owner id")
		return
	}

	tasks, err := store.GetTasksOfUser(ownerID, titleFilter)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithoutError(w, http.StatusOK, tasks, r.Header.Get("Content-type"))
}

func CreateTask(w http.ResponseWriter, r *http.Request, store stores.TaskStore) {
	var t model.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := store.CreateTask(&t); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithoutError(w, http.StatusCreated, t, r.Header.Get("Content-type"))
}

func UpdateTask(w http.ResponseWriter, r *http.Request, store stores.TaskStore) {
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

	if err := store.UpdateTask(&t); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithoutError(w, http.StatusOK, t, r.Header.Get("Content-type"))
}

func DeleteTask(w http.ResponseWriter, r *http.Request, store stores.TaskStore) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	t := model.Task{ID: id}
	if err := store.DeleteTask(&t); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithoutError(w, http.StatusCreated, map[string]string{"result": "success"}, r.Header.Get("Content-type"))
}
