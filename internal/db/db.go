package db

import (
	"context"
	"database/sql"
	"gopher/internal/env"
	"time"

	_ "github.com/lib/pq"
)

func New(addr string, maxIdleConns, maxOpenConns int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenConns)
	duration, err := env.GetDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(time.Duration(duration))
	db.SetMaxIdleConns(maxIdleConns)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
