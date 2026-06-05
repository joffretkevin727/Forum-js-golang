package model

import (
	"database/sql"
)

type LikeModel struct {
	DB *sql.DB
}

// VoteTarget applique un vote (-1, 0, 1) sur la table spécifiée ("liketopic" ou "likecomment") et sa colonne associée ("topicid" ou "commentid")
func (m *LikeModel) VoteTarget(table, column string, userID, targetID, val int) error {
	query := `INSERT INTO ` + table + ` (userid, ` + column + `, ` + "`like`" + `) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE ` + "`like`" + ` = ?`
	_, err := m.DB.Exec(query, userID, targetID, val, val)
	return err
}