package app

import (
	database "btct/database"
	"database/sql"
)

type App struct {
	db      *sql.DB
	Queries *database.Queries
}

// NewApp initializes a new app
func NewApp(db *sql.DB) *App {
	return &App{
		db:      db,
		Queries: database.New(db),
	}
}
