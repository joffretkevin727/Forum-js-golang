package model

import (
	"database/sql"
	structure "forum/Structure"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) GetByUsername(username string) (*structure.User, error) {

	var user structure.User

	query := `SELECT id, username, email, password_hash, role, is_banned, created_at FROM users WHERE username = ?`

	err := m.DB.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.IsBanned,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
