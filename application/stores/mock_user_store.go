package stores

import model "github.com/BogdanMelchenko/gorest-task/application/model"

func (db *MockDb) CreateUser(user *model.User) error {
	rets := db.Called(user)
	return rets.Error(0)
}

func (db *MockDb) UpdateUser(user *model.User) error {
	rets := db.Called(user)
	return rets.Error(0)
}
func (db *MockDb) DeleteUser(user *model.User) error {
	rets := db.Called(user)
	return rets.Error(0)
}
func (db *MockDb) GetUser(user *model.User) error {
	rets := db.Called(user)
	return rets.Error(0)
}
func (db *MockDb) GetUsers() ([]model.User, error) {
	rets := db.Called()
	return rets.Get(0).([]model.User), rets.Error(1)
}

func (db *MockDb) GetUsersFilteredByRole(role int) ([]model.User, error) {
	rets := db.Called()
	return rets.Get(0).([]model.User), rets.Error(1)
}
