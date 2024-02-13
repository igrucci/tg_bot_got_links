package postgres

import (
	"database/sql"
	"fmt"
)

type Storaage struct {
	db *sql.DB
}

func New(path string) (*Storaage, error) {
	db, err := sql.Open("postgres", path)
	if err != nil {
		return nil, fmt.Errorf("cant open database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("cant connect to database: %w", err)
	}
	return &Storaage{db: db}, nil
}
