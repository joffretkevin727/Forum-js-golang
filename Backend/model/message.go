package model

import (
	"database/sql"
	structure "forum/structure"
)

// ============================================================
// MESSAGE MODEL
// ============================================================
type MessageModel struct {
	DB *sql.DB
}

// Create ajoute un message en réponse à un sujet
func (m *MessageModel) Create(body string, topicID, authorID int) (int64, error) {
	query := `INSERT INTO messages (body, topic_id, author_id) VALUES (?, ?, ?)`
	res, err := m.DB.Exec(query, body, topicID, authorID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId() // Retourne l'ID du message créé
}

// GetByTopic récupère tous les messages d'un sujet (pour l'affichage du fil)
func (m *MessageModel) GetByTopic(topicID int) ([]structure.Message, error) {
	query := `SELECT id, body, topic_id, author_id, created_at FROM messages WHERE topic_id = ? ORDER BY created_at ASC`
	rows, err := m.DB.Query(query, topicID)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Libère la connexion après traitement

	var messages []structure.Message
	for rows.Next() {
		var msg structure.Message
		err := rows.Scan(&msg.ID, &msg.Body, &msg.TopicID, &msg.AuthorID, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

// Delete supprime un message via son ID
func (m *MessageModel) Delete(id int) error {
	query := `DELETE FROM messages WHERE id = ?`
	_, err := m.DB.Exec(query, id)
	return err // Retourne directement l'erreur
}
