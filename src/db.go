package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbDriver = "sqlite3"
	dbName   = "data/todo.db"
)

// Connect to the database
func connect() (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dbName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Get all items from the database
func GetItems(db *sql.DB) ([]Todo, error) {
	// Get all items from the database
	query := "SELECT * FROM todo"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows
	var items []Todo
	for rows.Next() {
		var item Todo
		if err := rows.Scan(&item.ID, &item.Title, &item.Completed); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	// Return the items
	return items, nil
}

// Add item
func AddItem(db *sql.DB, item *Todo) error {
	// Execute the insert statement
	_, err := db.Exec("INSERT INTO todo (title, completed) VALUES (?, ?)", item.Title, item.Completed)
	if err != nil {
		return err
	}

	return nil
}
