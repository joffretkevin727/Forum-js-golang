package model

import (
	"database/sql"
	"encoding/json"
	"forum/structure"
)

type TopicModel struct {
	DB *sql.DB
}

// GetAll récupère la totalité des sujets bruts pour laisser le front-end gérer le filtrage et la pagination
func (m *TopicModel) GetAll() ([]structure.Topic, error) {
	rows, err := m.DB.Query("SELECT id, title, body, status, author_id, created_at, pseudo, tags, like_count, dislike_count, isLike FROM topics")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []structure.Topic
	for rows.Next() {
		var t structure.Topic
		var tagsRaw []byte
		err := rows.Scan(&t.ID, &t.Title, &t.Body, &t.Status, &t.AuthorID, &t.CreatedAt, &t.Pseudo, &tagsRaw, &t.LikeCount, &t.DislikeCount, &t.IsLike)
		if err != nil {
			return nil, err
		}
		if tagsRaw != nil {
			json.Unmarshal(tagsRaw, &t.Tags)
		}
		list = append(list, t)
	}
	return list, nil
}

// Create insère un sujet avec ses tags encodés au format JSON
func (m *TopicModel) Create(title, body, pseudo string, authorID int, tags []string) (int64, error) {
	tagsJSON, err := json.Marshal(tags)
	if err != nil {
		return 0, err
	}

	query := "INSERT INTO topics (title, body, status, author_id, pseudo, tags) VALUES (?, ?, 'open', ?, ?, ?)"
	res, err := m.DB.Exec(query, title, body, authorID, pseudo, tagsJSON)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// Delete supprime définitivement un sujet par son ID
func (m *TopicModel) Delete(id int) error {
	_, err := m.DB.Exec("DELETE FROM topics WHERE id = ?", id)
	return err
}