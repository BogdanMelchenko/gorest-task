package model

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	Role int    `json:"role" xml:"role"`
}

func (u *User) GetUser(db *sql.DB) error {
	return db.QueryRow("SELECT name, role FROM users WHERE id=$1",
		u.ID).Scan(&u.Name, &u.Role)
}

func (u *User) UpdateUser(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE users SET name=$1, role=$2 WHERE id=$3",
			u.Name, u.Role, u.ID)

	return err
}

func (u *User) DeleteUser(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", u.ID)

	return err
}

func (u *User) CreateUser(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO users(name, role) VALUES($1, $2) RETURNING id",
		u.Name, u.Role).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(
		"SELECT id, name, role FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func GetUsersFilteredByRole(db *sql.DB, role int) ([]User, error) {
	rows, err := db.Query(
		"SELECT id, name, role FROM users WHERE role=$1", role)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
