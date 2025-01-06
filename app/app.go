package app

import (
	"context"
	"database/sql"
	"log"
)

type App struct {
	db      *sql.DB
	Queries *Queries
}

// NewApp initializes a new app
func NewApp(db *sql.DB) *App {
	return &App{
		db:      db,
		Queries: New(db),
	}
}

func (a *App) Init(ddl string) error {
	ctx := context.Background()
	_, err := a.Queries.db.ExecContext(ctx, ddl)
	if err != nil {
		log.Fatalf("An error occured creating database: %w", err)
	}

	return nil
}
