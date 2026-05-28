package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "forum/model"
)

type UserController struct {
	Model *model.UserModel
}

// GetUserHandler récupère un utilisateur par son username
func (c *UserController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	username := r.PathValue("username")
	if username == "" {
		http.Error(w, "Le paramètre username est requis", http.StatusBadRequest)
		return
	}
	u, err := c.Model.GetByUsername(username)
	if err != nil {
		http.Error(w, "Utilisateur introuvable", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(u)
}

// CreateUserHandler crée un nouvel utilisateur
func (c *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body struct {
		Username     string `json:"username"`
		Email        string `json:"email"`
		PasswordHash string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}
	id, err := c.Model.Create(body.Username, body.Email, body.PasswordHash)
	if err != nil {
		http.Error(w, "Erreur lors de la création", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

// BanUserHandler met à jour le statut de bannissement
func (c *UserController) BanUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	var body struct {
		IsBanned bool `json:"is_banned"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}
	if err := c.Model.UpdateBanStatus(id, body.IsBanned); err != nil {
		http.Error(w, "Erreur lors de la mise à jour", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
