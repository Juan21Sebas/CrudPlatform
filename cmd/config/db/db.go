package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgreSQLDB() (*sql.DB, error) {

	host := "localhost"
	port := "5432"
	dbname := "talentpitch"
	user := "juansebastiansanchez"
	password := "admin"

	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		host, port, dbname, user, password)

	db, err := sql.Open("postgres", connStr)
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

	// Creación tabla users
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE,
		image_path TEXT,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
	)`)
	if err != nil {
		fmt.Println("Error al crear la tabla users:", err)
		db.Close()
		return nil, err
	}

	// Creación tabla challenges
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS challenges (
		id TEXT PRIMARY KEY,
		title TEXT,
		description TEXT,
		difficulty INTEGER,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
	)`)
	if err != nil {
		fmt.Println("Error al crear la tabla challenges:", err)
		db.Close()
		return nil, err
	}

	// Creación tabla videos
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS videos (
		id TEXT PRIMARY KEY,
		title TEXT,
		description TEXT,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
	)`)
	if err != nil {
		fmt.Println("Error al crear la tabla videos:", err)
		db.Close()
		return nil, err
	}

	return db, nil
}
