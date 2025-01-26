package internals

type InsertTask struct {
	Name        string
	Description string
}

type Task struct {
	Id          int
	Name        string
	Description string
	Date        string
	Completed   bool
}
