package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	Owner_Id    int    `json:"owner_id"`
}

func (t *Task) GetTask(db *sql.DB) error {
	return db.QueryRow("SELECT title, description, done FROM tasks WHERE id=$1", t.ID).Scan(&t.Title, &t.Description, &t.Done)
}

func (t *Task) GetTaskOfUser(db *sql.DB) error {
	return db.QueryRow("SELECT title, description, done FROM tasks WHERE owner_id=$1 AND id=$2", t.Owner_Id, t.ID).Scan(&t.Title, &t.Description, &t.Done)
}

func GetTasksOfUser(db *sql.DB, owner_id int, titleFilter string) ([]Task, error) {
	filter := fmt.Sprintf("%%(%s)%%", titleFilter)
	rows, err := db.Query(
		"SELECt id, title, description, done FROM tasks WHERE owner_id=$1 and  title SIMILAR TO $2",
		owner_id, filter)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []Task{}

	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (t *Task) UpdateTask(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE tasks SET title=$1, description=$2, done=$3",
			t.Title, t.Description, t.Done)
	return err
}

func (t *Task) DeleteTask(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id=$1", t.ID)
	return err
}

func (t *Task) CreateTask(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO tasks(title, description, done) VALUES ($1, $2, $3) RETURNING id",
		t.Title, t.Description, t.Done).Scan(&t.ID)
	if err != nil {
		return err
	}

	return nil
}

func GetTasks(db *sql.DB) ([]Task, error) {
	rows, err := db.Query(
		"SELECt id, title, description,done, owner_id done FROM tasks")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []Task{}

	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Done, &t.Owner_Id); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil

}

func GetTasksFilteredByTitle(db *sql.DB, titleFilter string) ([]Task, error) {
	filter := fmt.Sprintf("%%(%s)%%", titleFilter)

	rows, err := db.Query(
		"SELECt id, title, description,done, owner_id done FROM tasks WHERE title SIMILAR TO $1", filter)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []Task{}

	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Done, &t.Owner_Id); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil

}
