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

// Delete item
func DeleteItem(db *sql.DB, id int) error {
	// Execute the delete statement
	_, err := db.Exec("DELETE FROM todo WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

// Update item
// Update item
func UpdateItem(db *sql.DB, id int, item *Todo) error {
	// Build the SQL query
	query := "UPDATE todo SET "
	args := []interface{}{}

	// Check if Title is specified
	if item.Title != nil {
		query += "title = ?, "
		args = append(args, item.Title)
	}

	// Check if Completed is specified
	if item.Completed != nil {
		query += "completed = ?, "
		args = append(args, *item.Completed)
	}

	// Remove the trailing comma if either Title or Completed is present
	if len(args) > 0 {
		query = query[:len(query)-2] // Remove the last ", "
	} else {
		// If neither Title nor Completed is specified, return nil without executing the query
		return nil
	}

	query += " WHERE id = ?"
	args = append(args, id)

	// Execute the update query
	_, err := db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
