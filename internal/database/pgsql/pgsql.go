package pgsql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PGSQL struct {
	db *sql.DB
}

func NewPGSQL(connStr string) (*PGSQL, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	return &PGSQL{db: db}, nil
}

func (p *PGSQL) Close() error {
	return p.db.Close()
}

func (p *PGSQL) Connect() error {
	if err := p.db.Ping(); err != nil {
		return fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}
	return nil
}
