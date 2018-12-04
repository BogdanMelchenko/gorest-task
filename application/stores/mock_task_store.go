package stores

import (
	model "github.com/BogdanMelchenko/gorest-task/application/model"
)

func (db *MockDb) GetTask(task *model.Task) error {
	rets := db.Called(task)
	return rets.Error(0)
}

func (db *MockDb) GetTaskOfUser(task *model.Task) error {
	rets := db.Called(task)
	return rets.Error(0)
}

func (db *MockDb) GetTasksOfUser(owner_id int, titleFilter string) ([]model.Task, error) {
	rets := db.Called()
	return rets.Get(0).([]model.Task), rets.Error(1)
}

func (db *MockDb) UpdateTask(task *model.Task) error {
	rets := db.Called(task)
	return rets.Error(0)
}

func (db *MockDb) DeleteTask(task *model.Task) error {
	rets := db.Called(task)
	return rets.Error(0)
}

func (db *MockDb) CreateTask(task *model.Task) error {
	rets := db.Called(task)
	return rets.Error(0)
}

func (db *MockDb) GetTasks() ([]model.Task, error) {
	rets := db.Called()
	return rets.Get(0).([]model.Task), rets.Error(1)
}

func (db *MockDb) GetTasksFilteredByTitle(titleFilter string) ([]model.Task, error) {
	rets := db.Called()
	return rets.Get(0).([]model.Task), rets.Error(1)
}
