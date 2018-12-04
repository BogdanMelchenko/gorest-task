package stores

import (
	model "github.com/BogdanMelchenko/gorest-task/application/model"
)

type UserStore interface {
	CreateUser(user *model.User) error
	DeleteUser(user *model.User) error
	UpdateUser(user *model.User) error
	GetUsersFilteredByRole(role int) ([]model.User, error)
	GetUsers() ([]model.User, error)
	GetUser(user *model.User) error
}

func (store *PostgresDbStore) CreateUser(user *model.User) error {
	err := store.Db.QueryRow(
		"INSERT INTO users(name, role) VALUES($1, $2) RETURNING id",
		user.Name, user.Role).Scan(&user.ID)

	if err != nil {
		return err
	}

	return nil
}

func (store *PostgresDbStore) GetUsersFilteredByRole(role int) ([]model.User, error) {
	rows, err := store.Db.Query(
		"SELECT id, name, role FROM users WHERE role=$1", role)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []model.User{}

	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (store *PostgresDbStore) GetUser(user *model.User) error {
	return store.Db.QueryRow("SELECT name, role FROM users WHERE id=$1",
		user.ID).Scan(&user.Name, &user.Role)
}

func (store *PostgresDbStore) UpdateUser(user *model.User) error {
	_, err :=
		store.Db.Exec("UPDATE users SET name=$1, role=$2 WHERE id=$3",
			user.Name, user.Role, user.ID)

	return err
}

func (store *PostgresDbStore) DeleteUser(user *model.User) error {
	_, err := store.Db.Exec("DELETE FROM users WHERE id=$1", user.ID)

	return err
}

func (store *PostgresDbStore) GetUsers() ([]model.User, error) {
	rows, err := store.Db.Query(
		"SELECT id, name, role FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []model.User{}

	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
