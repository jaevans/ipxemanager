package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func initializeDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./ipxe.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	err = createSchema(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create schema: %v", err)
	}

	return db, nil
}

func createSchema(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS systems (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		mac_address TEXT UNIQUE NOT NULL,
		dhcp_class TEXT,
		subnet TEXT,
		one_time_override BOOLEAN DEFAULT 0
	);

	CREATE TABLE IF NOT EXISTS configurations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL,
		content TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS system_configurations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		system_id INTEGER,
		configuration_id INTEGER,
		FOREIGN KEY (system_id) REFERENCES systems(id),
		FOREIGN KEY (configuration_id) REFERENCES configurations(id)
	);

	CREATE TABLE IF NOT EXISTS one_time_overrides (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		system_id INTEGER,
		override_content TEXT NOT NULL,
		FOREIGN KEY (system_id) REFERENCES systems(id)
	);

	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		token TEXT UNIQUE
	);
	`

	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to execute schema: %v", err)
	}

	return nil
}
