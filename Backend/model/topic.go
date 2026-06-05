package model

import (
	"database/sql"
	"encoding/json"
	structure "forum/structure"
)

type TopicModel struct {
	DB *sql.DB
}

func (modele *TopicModel) GetByID(id int) (*structure.Topic, error) {
	var topic structure.Topic
	var tagsRaw []byte
	query := "SELECT id, title, body, status, author_id, created_at, pseudo, tags, like_count, dislike_count, isLike FROM topics WHERE id = ?"
	err := modele.DB.QueryRow(query, id).Scan(
		&topic.ID,
		&topic.Title,
		&topic.Body,
		&topic.Status,
		&topic.AuthorID,
		&topic.CreatedAt,
		&topic.Pseudo,
		&tagsRaw,
		&topic.LikeCount,
		&topic.DislikeCount,
		&topic.IsLike,
	)
	if err != nil {
		return nil, err
	}
	if tagsRaw != nil {
		json.Unmarshal(tagsRaw, &topic.Tags)
	}
	return &topic, nil
}

func (modele *TopicModel) GetMany(limite int, offset int) ([]structure.Topic, error) {
	requete := "SELECT id, title, body, status, author_id, created_at, pseudo, tags, like_count, dislike_count, isLike FROM topics LIMIT ? OFFSET ?"
	lignes, err := modele.DB.Query(requete, limite, offset)
	if err != nil {
		return nil, err
	}
	defer lignes.Close()

	listeTopics := []structure.Topic{}
	for lignes.Next() {
		var topic structure.Topic
		var tagsRaw []byte
		err := lignes.Scan(
			&topic.ID,
			&topic.Title,
			&topic.Body,
			&topic.Status,
			&topic.AuthorID,
			&topic.CreatedAt,
			&topic.Pseudo,
			&tagsRaw,
			&topic.LikeCount,
			&topic.DislikeCount,
			&topic.IsLike,
		)
		if err != nil {
			return nil, err
		}
		if tagsRaw != nil {
			json.Unmarshal(tagsRaw, &topic.Tags)
		}
		listeTopics = append(listeTopics, topic)
	}
	return listeTopics, nil
}

// Single line comment: Marshals tags to JSON and inserts topic into MySQL.
func (modele *TopicModel) Create(title string, body string, authorID int, tags []string) (int64, error) {
    tagsJSON, err := json.Marshal(tags)
    if err != nil {
        return 0, err
    }

    query := "INSERT INTO topics (title, body, status, author_id, tags) VALUES (?, ?, 'open', ?, ?)"
    result, err := modele.DB.Exec(query, title, body, authorID, tagsJSON)
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}

func (modele *TopicModel) Update(id int, title string, body string, status string) error {
	query := "UPDATE topics SET title = ?, body = ?, status = ? WHERE id = ?"
	_, err := modele.DB.Exec(query, title, body, status, id)
	return err
}

func (modele *TopicModel) Delete(id int) error {
	query := "DELETE FROM topics WHERE id = ?"
	_, err := modele.DB.Exec(query, id)
	return err
}

func (modele *TopicModel) Count() (int, error) {
	var total int
	query := "SELECT COUNT(id) FROM topics"
	err := modele.DB.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (modele *TopicModel) GetByAuthor(authorID int) ([]structure.Topic, error) {
	query := "SELECT id, title, body, status, author_id, created_at, pseudo, tags, like_count, dislike_count, isLike FROM topics WHERE author_id = ?"
	rows, err := modele.DB.Query(query, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topics []structure.Topic
	for rows.Next() {
		var topic structure.Topic
		var tagsRaw []byte
		err := rows.Scan(
			&topic.ID,
			&topic.Title,
			&topic.Body,
			&topic.Status,
			&topic.AuthorID,
			&topic.CreatedAt,
			&topic.Pseudo,
			&tagsRaw,
			&topic.LikeCount,
			&topic.DislikeCount,
			&topic.IsLike,
		)
		if err != nil {
			return nil, err
		}
		if tagsRaw != nil {
			json.Unmarshal(tagsRaw, &topic.Tags)
		}
		topics = append(topics, topic)
	}
	return topics, nil
}
