package stores

import (
	"fmt"

	model "github.com/BogdanMelchenko/gorest-task/application/model"
)

type TaskStore interface {
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task) error
	DeleteTask(task *model.Task) error
	GetTask(task *model.Task) error
	GetTaskOfUser(task *model.Task) error
	GetTasks() ([]model.Task, error)
	GetTasksFilteredByTitle(titleFilter string) ([]model.Task, error)
	GetTasksOfUser(ownerID int, titleFilter string) ([]model.Task, error)
}

func (store *PostgresDbStore) GetTask(task *model.Task) error {
	return store.Db.QueryRow("SELECT title, description, done FROM tasks WHERE id=$1", task.ID).Scan(&task.Title, &task.Description, &task.Done)
}

func (store *PostgresDbStore) GetTaskOfUser(task *model.Task) error {
	return store.Db.QueryRow("SELECT title, description, done FROM tasks WHERE owner_id=$1 AND id=$2", task.OwnerID, task.ID).Scan(&task.Title, &task.Description, &task.Done)
}

func (store *PostgresDbStore) GetTasksOfUser(owner_id int, titleFilter string) ([]model.Task, error) {
	filter := fmt.Sprintf("%%(%s)%%", titleFilter)
	rows, err := store.Db.Query(
		"SELECt id, title, description, done FROM tasks WHERE owner_id=$1 and  title SIMILAR TO $2",
		owner_id, filter)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []model.Task{}

	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (store *PostgresDbStore) UpdateTask(task *model.Task) error {
	_, err :=
		store.Db.Exec("UPDATE tasks SET title=$1, description=$2, done=$3",
			task.Title, task.Description, task.Done)
	return err
}

func (store *PostgresDbStore) DeleteTask(task *model.Task) error {
	_, err := store.Db.Exec("DELETE FROM tasks WHERE id=$1", task.ID)
	return err
}

func (store *PostgresDbStore) CreateTask(task *model.Task) error {
	err := store.Db.QueryRow("INSERT INTO tasks(title, description, done) VALUES ($1, $2, $3) RETURNING id",
		task.Title, task.Description, task.Done).Scan(&task.ID)
	if err != nil {
		return err
	}

	return nil
}

func (store *PostgresDbStore) GetTasks() ([]model.Task, error) {
	rows, err := store.Db.Query(
		"SELECt id, title, description,done, owner_id done FROM tasks")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []model.Task{}

	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Done, &t.OwnerID); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (store *PostgresDbStore) GetTasksFilteredByTitle(titleFilter string) ([]model.Task, error) {
	filter := fmt.Sprintf("%%(%s)%%", titleFilter)

	rows, err := store.Db.Query(
		"SELECt id, title, description,done, owner_id done FROM tasks WHERE title SIMILAR TO $1", filter)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []model.Task{}

	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Done, &t.OwnerID); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}
