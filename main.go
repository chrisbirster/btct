package main

import (
	"context"
	"database/sql"
	"embed"
	"log"

	app "btct/app"
	migrations "btct/migrations"
	server "btct/server"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed dist/* dist/*/*
var staticFiles embed.FS

const DB_FILE = "btct.db"

func main() {
	// Connect to the SQLite database

	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		log.Fatalf("Failed to connect to SQLite: %v", err)
	}

	defer db.Close()

	err = migrations.ApplyMigrations(context.Background(), db)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	appInstance := app.NewApp(db)
	server.StartServer(appInstance, staticFiles)
}
