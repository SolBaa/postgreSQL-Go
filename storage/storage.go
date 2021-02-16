package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	// Pkg para conectar con postgres
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// NewPostrgesDB nueva ocnexion con PostreDB
func NewPostrgesDB() {
	once.Do(func() {
		var err error
		db, err := sql.Open("postgres", "postgres://postgres:12345@localhost:5432/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("Failed to conect: %v", err)
		}
		defer db.Close()

		if err = db.Ping(); err != nil {
			log.Fatalf("Failed to ping: %v", err)
		}
		fmt.Println("Conectado a Postgres")
	})
}

// Pool retorna una unica insatncia de db
func Pool() *sql.DB {
	return db
}
