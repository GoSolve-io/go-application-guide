package database

import (
	"fmt"

	"database/sql"

	// Needed for registering postgres drivers in sql package.
	_ "github.com/lib/pq"
)

// Adapter is a storage adapter for app.
type Adapter struct {
	db *sql.DB
}

// NewAdapter creates new db adapter.
func NewAdapter(dsn string) (*Adapter, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("opening postgres db: %w", err)
	}

	return &Adapter{
		db: db,
	}, nil
}

// Close closes underlying db client.
func (d *Adapter) Close() {
	d.db.Close()
}
