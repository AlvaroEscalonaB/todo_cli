package internals

import "time"

type InsertTask struct {
	Name        string
	Description string
}

type Task struct {
	Id          int
	Name        string
	Description string
	Date        time.Time
	Completed   bool
}
