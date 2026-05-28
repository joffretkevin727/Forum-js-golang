package model

import (
	"database/sql"
	structure "forum/structure"
)

type TopicModel struct {
	DB *sql.DB
}

// GetByID récupère un seul sujet par son identifiant unique
func (modele *TopicModel) GetByID(id int) (*structure.Topic, error) {
	var topic structure.Topic
	query := `SELECT id, title, body, status, author_id, created_at FROM topics WHERE id = ?`
	err := modele.DB.QueryRow(query, id).Scan(
		&topic.ID,
		&topic.Title,
		&topic.Body,
		&topic.Status,
		&topic.AuthorID,
		&topic.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &topic, nil
}

// GetMany récupère un groupe de topics en utilisant LIMIT et OFFSET pour la pagination
func (modele *TopicModel) GetMany(limite int, offset int) ([]structure.Topic, error) {
	requete := `SELECT id, title, body, status, author_id, created_at FROM topics LIMIT ? OFFSET ?`
	lignes, err := modele.DB.Query(requete, limite, offset)
	if err != nil {
		return nil, err
	}
	defer lignes.Close() // Libère la connexion MAMP une fois la fonction terminée

	listeTopics := []structure.Topic{}
	for lignes.Next() {
		var topic structure.Topic
		err := lignes.Scan(
			&topic.ID,
			&topic.Title,
			&topic.Body,
			&topic.Status,
			&topic.AuthorID,
			&topic.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		listeTopics = append(listeTopics, topic)
	}
	return listeTopics, nil
}

// Create insère un nouveau sujet dans la base de données
func (modele *TopicModel) Create(title string, body string, authorID int) (int64, error) {
	query := `INSERT INTO topics (title, body, status, author_id) VALUES (?, ?, 'open', ?)`
	result, err := modele.DB.Exec(query, title, body, authorID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId() // Retourne l'identifiant du topic créé
}

// Update modifie le titre, le contenu et le statut d'un sujet existant
func (modele *TopicModel) Update(id int, title string, body string, status string) error {
	query := `UPDATE topics SET title = ?, body = ?, status = ? WHERE id = ?`
	_, err := modele.DB.Exec(query, title, body, status, id)
	return err // Retourne directement l'erreur si elle existe
}

// Delete supprime définitivement un sujet de la base de données
func (modele *TopicModel) Delete(id int) error {
	query := `DELETE FROM topics WHERE id = ?`
	_, err := modele.DB.Exec(query, id)
	return err // Retourne directement l'erreur si elle existe
}

// Count récupère le nombre total de topics (indispensable pour calculer les pages)
func (modele *TopicModel) Count() (int, error) {
	var total int
	query := `SELECT COUNT(id) FROM topics`
	err := modele.DB.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

// GetByAuthor récupère la liste complète des sujets créés par un utilisateur spécifique
func (modele *TopicModel) GetByAuthor(authorID int) ([]structure.Topic, error) {
	query := `SELECT id, title, body, status, author_id, created_at FROM topics WHERE author_id = ?`
	rows, err := modele.DB.Query(query, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Libère la connexion après traitement

	var topics []structure.Topic
	for rows.Next() {
		var t structure.Topic
		err := rows.Scan(&t.ID, &t.Title, &t.Body, &t.Status, &t.AuthorID, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		topics = append(topics, t)
	}
	return topics, nil
}
