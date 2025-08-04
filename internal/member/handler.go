package member

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

type Handler struct {
    DB *sql.DB
}

func (h *Handler) CreateMember(w http.ResponseWriter, r *http.Request) {
    var m Member
    if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    err := h.DB.QueryRow(
        "INSERT INTO members (name, email, join_date, status) VALUES ($1, $2, $3, $4) RETURNING id",
        m.Name, m.Email, m.JoinDate, m.Status,
    ).Scan(&m.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(m)
}

func (h *Handler) GetMember(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var m Member
    err := h.DB.QueryRow("SELECT id, name, email, join_date, status FROM members WHERE id = $1", id).
        Scan(&m.ID, &m.Name, &m.Email, &m.JoinDate, &m.Status)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(m)
}
