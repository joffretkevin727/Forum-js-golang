package model

import (
	"database/sql"
	structure "forum/Structure"
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
