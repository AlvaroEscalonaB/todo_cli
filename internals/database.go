package internals

import (
	"database/sql"
	"fmt"
	"os"
)

const TableName string = "tasks"
const dbUrl string = "db/db.sqlite"

const createDatabaseInitialString string = `
  CREATE TABLE IF NOT EXISTS tasks (
  id INTEGER NOT NULL PRIMARY KEY,
	name STRING NOT NULL,
  description TEXT,
  date DATETIME NOT NULL,
	completed BOOL NOT NULL
);`

func SetupFileDatabase() error {
	folderPath := "./db"
	filePath := "./db/db.sqlite"

	// Ensure the folder exists (creates it if necessary)
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create folder: %v\n", err)
		return err
	}

	// Ensure the file exists (creates it if necessary)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Failed to create or open file: %v\n", err)
		return err
	}
	defer file.Close()
	return nil
}

func CreateDatabase() error {
	err := SetupFileDatabase()
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", dbUrl)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(createDatabaseInitialString)

	if err != nil {
		fmt.Println("Cannot create the database")
		return err
	}
	return nil
}

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}
