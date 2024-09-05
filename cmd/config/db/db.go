package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteDB() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./talentpitch.db")
	if err != nil {
		fmt.Println("Error al abrir la base de datos:", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Error de conexión a la base de datos:", err)
		db.Close()
		return nil, err
	}

	tables := []string{
		"videos",
		"challenges",
		"users",
	}

	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table))
		if err != nil {
			return nil, fmt.Errorf("error dropping table %s: %w", table, err)
		}
	}

	// Creacion tabla users
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE,
		image_path TEXT,
		created_at TEXT,
		updated_at TEXT
	)`)
	if err != nil {
		fmt.Println("Error al crear la tabla users:", err)
		db.Close()
		return nil, err
	}

	// Creacion tabla challenges
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS challenges (
		id TEXT PRIMARY KEY,
		title TEXT,
		description TEXT,
		difficulty INTEGER,
		created_at TEXT,
		updated_at TEXT
	)`)
	if err != nil {
		fmt.Println("Error al crear la tabla challenges:", err)
		db.Close()
		return nil, err
	}

	// Creacion tabla videos
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS videos (
		id TEXT PRIMARY KEY,
		title TEXT,
		description TEXT,
		created_at TEXT,
		updated_at TEXT
	)`)
	if err != nil {
		fmt.Println("Error al crear la tabla videos:", err)
		db.Close()
		return nil, err
	}

	return db, nil
}
