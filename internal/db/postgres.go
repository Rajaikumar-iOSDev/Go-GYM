package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

func NewSupabaseDB() *sql.DB {
    connStr := "postgresql://postgres:Ziffity@123@db.dzomirjppizanvfqprni.supabase.co:5432/postgres"
    if connStr == "" {
        log.Fatal("SUPABASE_DB_URL environment variable is not set")
    }

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Failed to connect to Supabase:", err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal("Supabase database unreachable:", err)
    }
    return db
}
