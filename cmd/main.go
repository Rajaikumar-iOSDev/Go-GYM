package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/Rajaikumar-iOSDev/go-gym/internal/db"
    "github.com/Rajaikumar-iOSDev/go-gym/internal/member"
)

func main() {
    connStr := os.Getenv("POSTGRES_CONN")
    if connStr == "" {
        connStr = "host=localhost port=5432 user=postgres password=postgres dbname=gym sslmode=disable"
    }
    database := db.NewSupabaseDB()
    memberHandler := &member.Handler{DB: database}

    r := mux.NewRouter()
    r.HandleFunc("/members", memberHandler.CreateMember).Methods("POST")
    r.HandleFunc("/members/{id}", memberHandler.GetMember).Methods("GET")

    fmt.Println("Server running at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
