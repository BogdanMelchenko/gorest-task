package handlertest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/BogdanMelchenko/gorest-task/application/model"

	"github.com/BogdanMelchenko/gorest-task/application/handlers"
	"github.com/BogdanMelchenko/gorest-task/application/stores"
)

func TestGetTasksHandler(t *testing.T) {
	var db stores.MockDb
	tmptask := model.Task{ID: 1, Title: "ttask", OwnerID: 1, Done: false, Description: "ddescr"}
	db.On("GetTasks").Return([]model.Task{
		tmptask}, nil).Once()

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasks(w, r, &db)
	})

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := tmptask
	b := []model.Task{}

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

func TestGetTaskHandler(t *testing.T) {
	var db stores.MockDb
	tmptask := &model.Task{ID: 1, Title: "", OwnerID: 0, Done: false, Description: ""}
	db.On("GetTask", tmptask).Return(nil).Once()

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTask(w, r, &db)
	})

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := tmptask
	b := &model.Task{}

	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b

	if !assert.ObjectsAreEqual(expected, actual) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
	db.AssertExpectations(t)
}

func TestCreateTaskHandler(t *testing.T) {
	var db stores.MockDb
	tmpuser := &model.Task{ID: 0, Title: "ttask", OwnerID: 2, Description: "ddescr", Done: false}
	db.On("CreateTask", tmpuser).Return(nil)

	form, _ := json.Marshal(tmpuser)

	req, err := http.NewRequest("POST", "", bytes.NewBuffer(form))

	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTask(w, r, &db)
	})

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	db.AssertExpectations(t)
}
