package model

import (
	"database/sql"
)

// ============================================================
// MESSAGE VOTE MODEL
// ============================================================
type MessageVoteModel struct {
	DB *sql.DB
}

// Vote ajoute ou met à jour le vote (+1 ou -1) d'un utilisateur sur un message
func (modele *MessageVoteModel) Vote(userID, messageID, vote int) error {
	query := `INSERT INTO message_votes (user_id, message_id, vote) 
			  VALUES (?, ?, ?) 
			  ON DUPLICATE KEY UPDATE vote = ?`
	_, err := modele.DB.Exec(query, userID, messageID, vote, vote)
	return err // Gère l'insertion ou la modification transparente
}

// GetScore calcule la somme des votes (+1 et -1) pour un message donné
func (modele *MessageVoteModel) GetScore(messageID int) (int, error) {
	var score sql.NullInt64
	query := `SELECT SUM(vote) FROM message_votes WHERE message_id = ?`
	err := modele.DB.QueryRow(query, messageID).Scan(&score)
	if err != nil {
		return 0, err
	}
	if !score.Valid {
		return 0, nil // Retourne 0 si aucun vote n'est enregistré
	}
	return int(score.Int64), nil
}
