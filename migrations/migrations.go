package migrations

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed *.sql
var migrationFiles embed.FS

// Migration represents a single migration file
type Migration struct {
	Count       int
	Description string
	SQL         string
	Filename    string
}

// GetAllMigrations retrieves and parses all migration files
func GetAllMigrations() ([]Migration, error) {
	entries, err := migrationFiles.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded migrations: %w", err)
	}

	var migrations []Migration
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}

		// Extract migration count and description from filename
		parts := strings.SplitN(entry.Name(), "_", 2)
		if len(parts) != 2 {
			log.Printf("Skipping invalid migration filename: %s", entry.Name())
			continue
		}

		count, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Printf("Skipping migration with invalid count: %s", entry.Name())
			continue
		}

		description := strings.TrimSuffix(parts[1], ".sql")

		// Read migration SQL
		sqlBytes, err := migrationFiles.ReadFile(entry.Name())
		if err != nil {
			return nil, fmt.Errorf("failed to read migration file %s: %w", entry.Name(), err)
		}

		migrations = append(migrations, Migration{
			Count:       count,
			Description: description,
			SQL:         string(sqlBytes),
			Filename:    entry.Name(),
		})
	}

	// Sort migrations by Count
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Count < migrations[j].Count
	})

	return migrations, nil
}

// GetAppliedMigrations retrieves applied migration counts from the database
func GetAppliedMigrations(ctx context.Context, db *sql.DB) (map[int]bool, error) {
	rows, err := db.QueryContext(ctx, "SELECT count FROM migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to query migrations table: %w", err)
	}
	defer rows.Close()

	applied := make(map[int]bool)
	for rows.Next() {
		var count int
		if err := rows.Scan(&count); err != nil {
			return nil, fmt.Errorf("failed to scan migration count: %w", err)
		}
		applied[count] = true
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating migration rows: %w", err)
	}

	return applied, nil
}

// ApplyMigrations applies all pending migrations in order
func ApplyMigrations(ctx context.Context, db *sql.DB) error {
	// Ensure migrations table exists
	_, err := db.ExecContext(ctx, `
        CREATE TABLE IF NOT EXISTS migrations (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            count INTEGER UNIQUE NOT NULL,
            description TEXT NOT NULL,
            applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	allMigrations, err := GetAllMigrations()
	if err != nil {
		return fmt.Errorf("failed to get all migrations: %w", err)
	}

	appliedMigrations, err := GetAppliedMigrations(ctx, db)
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}

	for _, m := range allMigrations {
		if appliedMigrations[m.Count] {
			log.Printf("Migration %d already applied: %s", m.Count, m.Description)
			continue
		}

		log.Printf("Applying migration %d: %s", m.Count, m.Description)

		// Begin transaction
		tx, err := db.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}

		// Execute migration SQL
		_, err = tx.ExecContext(ctx, m.SQL)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to execute migration %d (%s): %w", m.Count, m.Filename, err)
		}

		// Insert migration record
		_, err = tx.ExecContext(ctx, `
            INSERT INTO migrations (count, description)
            VALUES (?, ?)
        `, m.Count, m.Description)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to record migration %d: %w", m.Count, err)
		}

		// Commit transaction
		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit migration %d: %w", m.Count, err)
		}

		log.Printf("Successfully applied migration %d: %s", m.Count, m.Description)
	}

	return nil
}
