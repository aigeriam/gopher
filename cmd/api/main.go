package main

import (
	"gopher/internal/db"
	"gopher/internal/env"
	"gopher/internal/store"
	"log"
)

const version = "0.01"

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/gopher?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_Idle_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_Idle_Time", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}	

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxIdleConns,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Println("database connection established")
	store := store.NewStorage(db)
	app := application{
		config: cfg,
		store:  store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
