package controller

import (
	"encoding/json"
	"forum/model"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Model *model.UserModel
}

// LoginHandler authentifie un utilisateur en vérifiant son email, son statut et son mot de passe haché.
func (c *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, `{"message":"Format JSON invalide"}`, http.StatusBadRequest)
		return
	}
	u, err := c.Model.GetByEmail(body.Email)
	if err != nil || u.IsBanned || bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(body.Password)) != nil {
		http.Error(w, `{"message":"Identifiants invalides ou compte banni"}`, http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"user": u})
}

// CreateUserHandler hache le mot de passe reçu et insère un nouvel utilisateur dans la base de données.
func (c *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	// Single line comment: calls model create method with parsed parameters.
	id, err := c.Model.Create(body.Username, body.Email, string(hash))
	if err != nil {
		http.Error(w, "Erreur lors de la création", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

// BanUserHandler modifie le statut de bannissement d'un utilisateur ciblé par son identifiant.
func (c *UserController) BanUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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

// DeleteUserHandler supprime définitivement un compte utilisateur de la base de données.
func (c *UserController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	if err := c.Model.Delete(id); err != nil {
		http.Error(w, "Erreur de suppression", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}