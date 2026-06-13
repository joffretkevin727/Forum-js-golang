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
	rows, err := m.DB.Query(`SELECT id, body, topic_id, author_id, created_at FROM messages`)
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

// GetAllById récupère tous les commentaires associés à un ID spécifique avec le pseudo de l'auteur.
func (m *CommentModel) GetAllById(topicID int) ([]structure.Comment, error) {
	// La requête SQL récupère les champs du message + le username de la table users
	query := `
        SELECT m.id, m.body, m.topic_id, m.author_id, u.username, m.created_at 
        FROM messages m
        INNER JOIN users u ON m.author_id = u.id
        WHERE m.topic_id = ?`

	rows, err := m.DB.Query(query, topicID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []structure.Comment
	for rows.Next() {
		var c structure.Comment
		// On ajoute &c.Pseudo dans le Scan pour récupérer le nom de l'auteur
		if err := rows.Scan(&c.ID, &c.Body, &c.TopicID, &c.AuthorID, &c.Pseudo, &c.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
