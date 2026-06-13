package structure

import "time"

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // Masqué dans les réponses JSON pour la sécurité
	IsBanned     bool      `json:"is_banned"`
	CreatedAt    time.Time `json:"created_at"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Topic struct {
	ID           int      `json:"id"`
	Title        string   `json:"title"`
	Body         string   `json:"text"`
	Status       string   `json:"status"`
	AuthorID     int      `json:"author_id"`
	CreatedAt    string   `json:"date"`
	Pseudo       string   `json:"pseudo"`
	Tags         []string `json:"tags"`
	LikeCount    int      `json:"upVotes"`
	DislikeCount int      `json:"downVotes"`
	IsLike       bool     `json:"isLike"`
}

type TopicTag struct {
	TopicID int `json:"topic_id"`
	TagID   int `json:"tag_id"`
}

type Comment struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	TopicID   int       `json:"topic_id"`
	AuthorID  int       `json:"author_id"`
	Pseudo    string    `json:"pseudo"`
	CreatedAt time.Time `json:"created_at"`
}

type LikeTopic struct {
	ID      int `json:"id"`
	UserID  int `json:"user_id"`
	TopicID int `json:"topic_id"`
	Like    int `json:"like"`
}

type LikeComment struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	CommentID int `json:"comment_id"`
	TopicID   int `json:"topic_id"`
	Like      int `json:"like"`
}
