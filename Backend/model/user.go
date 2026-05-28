package model

import (
	"database/sql"
	structure "forum/structure"
)

// ============================================================
// USER MODEL
// ============================================================
type UserModel struct {
	DB *sql.DB
}

// Create insère un nouvel utilisateur
func (m *UserModel) Create(username, email, passwordHash string) (int64, error) {
	query := `INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)`
	res, err := m.DB.Exec(query, username, email, passwordHash)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId() // Retourne l'ID de l'utilisateur créé
}

// GetByID récupère un utilisateur par son identifiant
func (m *UserModel) GetByID(id int) (*structure.User, error) {
	var u structure.User
	query := `SELECT id, username, email, password_hash, role, is_banned, created_at FROM users WHERE id = ?`
	err := m.DB.QueryRow(query, id).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.Role, &u.IsBanned, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// GetByUsername récupère un utilisateur par son nom d'utilisateur
func (m *UserModel) GetByUsername(username string) (*structure.User, error) {
	var u structure.User
	query := `SELECT id, username, email, password_hash, role, is_banned, created_at FROM users WHERE username = ?`
	err := m.DB.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.Role, &u.IsBanned, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// GetByEmail récupère un utilisateur par son email (utile pour la connexion)
func (m *UserModel) GetByEmail(email string) (*structure.User, error) {
	var u structure.User
	query := `SELECT id, username, email, password_hash, role, is_banned, created_at FROM users WHERE email = ?`
	err := m.DB.QueryRow(query, email).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.Role, &u.IsBanned, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// UpdateBanStatus modifie le statut de bannissement d'un utilisateur
func (m *UserModel) UpdateBanStatus(id int, isBanned bool) error {
	query := `UPDATE users SET is_banned = ? WHERE id = ?`
	_, err := m.DB.Exec(query, isBanned, id)
	return err // Retourne directement l'erreur
}
