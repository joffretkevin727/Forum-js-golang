package model

import (
	"database/sql"
	"forum/structure"
)

type UserModel struct {
	DB *sql.DB
}

// Create insère un nouvel utilisateur avec son mot de passe déjà haché
func (m *UserModel) Create(username, email, passwordHash string) (int64, error) {
	res, err := m.DB.Exec(`INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)`, username, email, passwordHash)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// GetByEmail récupère un utilisateur par son email pour la connexion
func (m *UserModel) GetByEmail(email string) (*structure.User, error) {
	var u structure.User
	query := `SELECT id, username, email, password_hash, is_banned, created_at FROM users WHERE email = ?`
	err := m.DB.QueryRow(query, email).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.IsBanned, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// Delete supprime définitivement un utilisateur par son ID
func (m *UserModel) Delete(id int) error {
	_, err := m.DB.Exec(`DELETE FROM users WHERE id = ?`, id)
	return err
}

// UpdateBanStatus modifie le statut de bannissement d'un utilisateur (0 ou 1)
func (m *UserModel) UpdateBanStatus(id int, isBanned bool) error {
	_, err := m.DB.Exec(`UPDATE users SET is_banned = ? WHERE id = ?`, isBanned, id)
	return err
}