package internals

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type TaskRepository struct {
	Db *sql.DB
}

func (taskRepository TaskRepository) NewTask(task InsertTask) (int, error) {
	result, err := taskRepository.Db.Exec(
		"INSERT INTO tasks (name, description, date) VALUES(?, ?, ?);",
		TableName, task.Name, task.Description, time.Now(),
	)
	if err != nil {
		return 0, err
	}

	var id int64

	if _, err := result.LastInsertId(); err != nil {
		return 0, err
	}

	return int(id), nil
}

func (taskRepository TaskRepository) ListTask() ([]Task, error) {
	var tasks []Task

	query := fmt.Sprintf("SELECT * FROM %s", TableName)
	results, err := taskRepository.Db.Query(query)

	if err != nil {
		return tasks, err
	}
	defer results.Close()

	for results.Next() {
		var task Task

		results.Scan(&task.id, &task.name, &task.description, &task.date)
		tasks = append(tasks, task)
	}
	defer taskRepository.Db.Close()
	return tasks, nil
}
