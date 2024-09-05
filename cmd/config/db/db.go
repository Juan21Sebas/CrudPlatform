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
		"program_participants",
		"programs",
		"companies",
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
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		difficulty INTEGER,
		user_id INTEGER,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`)
	if err != nil {
		fmt.Println("Error al crear la tabla challenges:", err)
		db.Close()
		return nil, err
	}

	// Creacion tabla companies
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS companies (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		image_path TEXT,
		location TEXT,
		industry TEXT,
		user_id INTEGER,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`)
	if err != nil {
		fmt.Println("Error al crear la tabla companies:", err)
		db.Close()
		return nil, err
	}

	// Creacion tabla programs
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS programs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		start_date DATE,
		end_date DATE,
		user_id INTEGER,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`)
	if err != nil {
		fmt.Println("Error al crear la tabla programs:", err)
		db.Close()
		return nil, err
	}

	// Creacion tabla program_participants
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS program_participants (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		program_id INTEGER,
		entity_type TEXT,
		entity_id INTEGER,
		FOREIGN KEY (program_id) REFERENCES programs(id)
	)`)
	if err != nil {
		fmt.Println("Error al crear la tabla program_participants:", err)
		db.Close()
		return nil, err
	}

	return db, nil
}
