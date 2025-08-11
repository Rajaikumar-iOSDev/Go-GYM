// package main

// import (
//     "fmt"
//     "log"
//     "net/http"
//     "os"

//     "github.com/gorilla/mux"
//     "github.com/Rajaikumar-iOSDev/go-gym/internal/db"
//     "github.com/Rajaikumar-iOSDev/go-gym/internal/member"
// )

// func main() {
//     connStr := os.Getenv("POSTGRES_CONN")
//     if connStr == "" {
//         connStr = "host=localhost port=5432 user=postgres password=postgres dbname=gym sslmode=disable"
//     }
//     database := db.NewSupabaseDB()
//     memberHandler := &member.Handler{DB: database}

//     r := mux.NewRouter()
//     r.HandleFunc("/members", memberHandler.CreateMember).Methods("POST")
//     r.HandleFunc("/members/{id}", memberHandler.GetMember).Methods("GET")

//     fmt.Println("Server running at :8080")
//     log.Fatal(http.ListenAndServe(":8080", r))
// }
package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func main() {
    projectID := "dzomirjppizanvfqprni"
    table := "members"
    apiKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImR6b21pcmpwcGl6YW52ZnFwcm5pIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NTQ0ODg0NzUsImV4cCI6MjA3MDA2NDQ3NX0.IDU9b9dH46AJso8Jeg2qC0ToqiRqj78GB2TyUJOXQn4"

    url := fmt.Sprintf("https://%s.supabase.co/rest/v1/%s", projectID, table)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal("Request creation failed:", err)
    }

    req.Header.Set("apikey", apiKey)
    req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Prefer", "return=representation") // ensures full row data is returned

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Request failed:", err)
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    fmt.Println("Response:", string(body))
}
