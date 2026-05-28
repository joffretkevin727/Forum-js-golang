package model

import (
	"database/sql"
	structure "forum/structure"
)

// ============================================================
// TAG MODEL
// ============================================================
type TagModel struct {
	DB *sql.DB
}

// Create insère un nouveau tag unique
func (m *TagModel) Create(name string) (int64, error) {
	query := `INSERT INTO tags (name) VALUES (?)`
	res, err := m.DB.Exec(query, name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId() // Retourne l'ID du tag créé
}

// GetAll récupère tous les tags disponibles
func (m *TagModel) GetAll() ([]structure.Tag, error) {
	query := `SELECT id, name FROM tags`
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Libère la connexion après traitement

	var tags []structure.Tag
	for rows.Next() {
		var t structure.Tag
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}
	return tags, nil
}

// AttachToTopic associe un tag à un sujet (remplit topic_tags)
func (m *TagModel) AttachToTopic(topicID, tagID int) error {
	query := `INSERT INTO topic_tags (topic_id, tag_id) VALUES (?, ?)`
	_, err := m.DB.Exec(query, topicID, tagID)
	return err // Retourne directement l'erreur
}

// GetByTopic récupère tous les tags liés à un sujet spécifique
func (m *TagModel) GetByTopic(topicID int) ([]structure.Tag, error) {
	query := `SELECT t.id, t.name FROM tags t 
			  JOIN topic_tags tt ON t.id = tt.tag_id 
			  WHERE tt.topic_id = ?`
	rows, err := m.DB.Query(query, topicID)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Libère la connexion après traitement

	var tags []structure.Tag
	for rows.Next() {
		var t structure.Tag
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}
	return tags, nil
}
