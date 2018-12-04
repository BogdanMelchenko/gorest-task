package handlertest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/BogdanMelchenko/gorest-task/application/model"

	"github.com/BogdanMelchenko/gorest-task/application/handlers"
	"github.com/BogdanMelchenko/gorest-task/application/stores"
)

func TestGetUsersHandler(t *testing.T) {
	var db stores.MockDb
	db.On("GetUsers").Return([]model.User{
		{ID: 1, Name: "Igar", Role: 1}}, nil).Once()

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUsers(w, r, &db)
	})

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := model.User{ID: 1, Name: "Igar", Role: 1}
	b := []model.User{}

	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if !assert.True(t, (actual == expected)) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
	db.AssertExpectations(t)
}

func TestCreateUserHandler(t *testing.T) {
	var db stores.MockDb
	tmpuser := &model.User{ID: 0, Name: "Ivanc", Role: 2}
	db.On("CreateUser", tmpuser).Return(nil)

	form, _ := json.Marshal(tmpuser)

	req, err := http.NewRequest("POST", "", bytes.NewBuffer(form))

	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUser(w, r, &db)
	})

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	db.AssertExpectations(t)
}
