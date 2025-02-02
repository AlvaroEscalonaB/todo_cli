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
		"INSERT INTO tasks (name, description, date, completed) VALUES(?, ?, ?, ?);",
		task.Name, task.Description, time.Now(), false,
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

func (taskRepository TaskRepository) ListTask(completed bool, uncompleted bool) ([]Task, error) {
	var tasks []Task
	var query string

	query = fmt.Sprintf("SELECT id, name, description, date, completed FROM %s", TableName)

	if completed {
		query = fmt.Sprintf("SELECT id, name, description, date, completed FROM %s WHERE completed = TRUE", TableName)
	}

	if uncompleted {
		query = fmt.Sprintf("SELECT id, name, description, date, completed FROM %s WHERE completed != TRUE", TableName)
	}

	results, err := taskRepository.Db.Query(query)

	if err != nil {
		return tasks, err
	}
	defer results.Close()

	for results.Next() {
		var task Task

		err := results.Scan(&task.Id, &task.Name, &task.Description, &task.Date, &task.Completed)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	defer taskRepository.Db.Close()
	return tasks, nil
}

func (taskRepository TaskRepository) CompleteTask(id int) (Task, error) {
	var task Task
	taskResult := taskRepository.Db.QueryRow("SELECT id, name, description, date, completed FROM tasks WHERE id = ?", id)

	if err := taskResult.Scan(&task.Id, &task.Name, &task.Description, &task.Date, &task.Completed); err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return task, fmt.Errorf("no task found with id: %d", id)
		}
		return task, err
	}

	if task.Completed {
		return task, nil
	}

	_, err := taskRepository.Db.Exec("UPDATE tasks SET completed = TRUE WHERE id = ?", id)

	if err != nil {
		fmt.Println("Cannot update the task")
		return task, err
	}

	task.Completed = true

	return task, nil
}
