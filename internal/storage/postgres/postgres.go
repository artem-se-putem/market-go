package postgres

import (
	"database/sql"
	"fmt"

	"market-go/internal/config"
)

type Storage struct {
	db *sql.DB
}

func New(cfg *config.Postgres) (*Storage, error){
	const op = "storage.sqlite.New"

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
		cfg.User, cfg.Password, cfg.Dbname, cfg.Sslmode, cfg.Host, cfg.Port)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
       return nil, fmt.Errorf("%s: failed to open connection: %w", op, err)
    }
    defer db.Close()

	return &Storage{db: db}, nil
}