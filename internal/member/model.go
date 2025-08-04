package member

type Member struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    JoinDate string `json:"join_date"`
    Status   string `json:"status"`
}
