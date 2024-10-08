package main

import (
	"CrudPlatform/cmd/config/db"
	"CrudPlatform/internal/adapters/handlers/http"
	"log"
)

func main() {
	dbInstance, err := db.NewPostgreSQLDB()
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	http.RunServer(dbInstance)
}
