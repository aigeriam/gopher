package main

import (
	"gopher/internal/env"
	"gopher/internal/store"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://user:adminpassword@localhost/gopher?sslmode=disabled"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_Idle_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_Idle_Time", "15min"),
		},
	}
	store := store.NewStorage(nil)
	app := application{
		config: cfg,
		store:  store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
