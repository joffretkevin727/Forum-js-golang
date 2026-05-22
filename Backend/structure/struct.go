package structure

import "time"

// User définit les champs de ton utilisateur (mappé sur ton SQL)
type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	IsBanned     bool      `json:"is_banned"`
	CreatedAt    time.Time `json:"created_at"`
}

type Topic struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Status    string    `json:"status"` // 'open', 'closed', 'archived'
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}
