package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

func NewPostgresDB(connStr string) *sql.DB {
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal("Database unreachable:", err)
    }
    return db
}
