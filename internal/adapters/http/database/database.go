package database

import (
	"fmt"

	"github.com/boltdb/bolt"
)

// Adapter is a storage adapter for app.
type Adapter struct {
	bdb *bolt.DB
}

// NewAdapter creates new db adapter.
func NewAdapter(path string) (*Adapter, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("opening bolt db file '%s': %w", path, err)
	}

	return &Adapter{
		bdb: db,
	}, nil
}

// Close closes underlying db client.
func (d *Adapter) Close() {
	d.bdb.Close()
}
