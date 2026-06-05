package model

import (
	"database/sql"
	"forum/structure"
)

type CommentModel struct {
	DB *sql.DB
}

// Create insère un commentaire en base de données
func (m *CommentModel) Create(body string, topicID, authorID int) error {
	_, err := m.DB.Exec(`INSERT INTO comments (body, topic_id, author_id) VALUES (?, ?, ?)`, body, topicID, authorID)
	return err
}

// GetAll récupère la totalité des commentaires pour laisser le front-end filtrer par topic_id
func (m *CommentModel) GetAll() ([]structure.Comment, error) {
	rows, err := m.DB.Query(`SELECT id, body, topic_id, author_id, created_at FROM comments`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []structure.Comment
	for rows.Next() {
		var c structure.Comment
		if err := rows.Scan(&c.ID, &c.Body, &c.TopicID, &c.AuthorID, &c.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}