package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	conf := "postgresql://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@localhost:5432/go-chat?sslmode=disable"
	db, err := sql.Open("postgres", conf)
	if err != nil {
		return nil, fmt.Errorf("cannot open db: %v\n", err)
	}
	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	err := d.db.Close()
	if err != nil {
		return fmt.Errorf("cannot close db connection: %v\n", err)
	}

	return nil
}

// TODO: should return error?
func (d *Database) GetDB() (*sql.DB, error) {
	db := d.db

	return db, nil
}
