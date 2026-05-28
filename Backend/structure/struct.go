package structure

import "time"

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	IsBanned     bool      `json:"is_banned"`
	CreatedAt    time.Time `json:"created_at"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Topic struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Status    string    `json:"status"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}

type TopicTag struct {
	TopicID int `json:"topic_id"`
	TagID   int `json:"tag_id"`
}

type Message struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	TopicID   int       `json:"topic_id"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}

type MessageVote struct {
	UserID    int `json:"user_id"`
	MessageID int `json:"message_id"`
	Vote      int `json:"vote"`
}
