package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	model "github.com/BogdanMelchenko/gorest-task/application/model"
	stores "github.com/BogdanMelchenko/gorest-task/application/stores"
	util "github.com/BogdanMelchenko/gorest-task/application/util"
	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request, store stores.UserStore) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	u := model.User{ID: id}
	if err := store.GetUser(&u); err != nil {
		switch err {
		case sql.ErrNoRows:
			util.RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	util.RespondWithoutError(w, http.StatusOK, u, r.Header.Get("Content-type"))
}

func GetUsers(w http.ResponseWriter, r *http.Request, store stores.UserStore) {

	v := r.URL.Query()

	roleFilterString := v.Get("role")
	roleFilter, err := strconv.Atoi(roleFilterString)
	if roleFilterString != "" && err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "role parameter format error")
		return
	}

	var users []model.User
	if roleFilterString != "" && err == nil {
		users, err = store.GetUsersFilteredByRole(roleFilter)
	} else {
		users, err = store.GetUsers()
	}
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithoutError(w, http.StatusOK, users, r.Header.Get("Content-type"))
}

func CreateUser(w http.ResponseWriter, r *http.Request, store stores.UserStore) {
	var u model.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := store.CreateUser(&u); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithoutError(w, http.StatusCreated, u, r.Header.Get("Content-type"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request, store stores.UserStore) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var u model.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	u.ID = id

	if err := store.UpdateUser(&u); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithoutError(w, http.StatusOK, u, r.Header.Get("Content-type"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request, store stores.UserStore) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	u := model.User{ID: id}
	if err := store.DeleteUser(&u); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithoutError(w, http.StatusOK, map[string]string{"result": "success"}, r.Header.Get("Content-type"))
}
