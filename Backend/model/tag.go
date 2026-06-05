package model

import (
	"database/sql"
	"forum/structure"
)

type TagModel struct {
	DB *sql.DB
}

// GetAll récupère tous les tags disponibles sur la plateforme
func (m *TagModel) GetAll() ([]structure.Tag, error) {
	rows, err := m.DB.Query(`SELECT id, name FROM tags`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

// AttachToTopic associe un tag à un sujet dans la table de liaison topic_tags
func (m *TagModel) AttachToTopic(topicID, tagID int) error {
	_, err := m.DB.Exec(`INSERT INTO topic_tags (topic_id, tag_id) VALUES (?, ?)`, topicID, tagID)
	return err
}