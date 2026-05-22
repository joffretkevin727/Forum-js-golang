package controller

import (
	"encoding/json"
	"net/http"

	model "forum/model"
)

type UserController struct {
	Model *model.UserModel
}

func (c *UserController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := r.PathValue("username")
	if username == "" {
		http.Error(w, "Le paramètre username est requis", http.StatusBadRequest)
		return
	}

	// Appel au Modèle
	u, err := c.Model.GetByUsername(username)
	if err != nil {
		http.Error(w, "Utilisateur introuvable", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(u)
}
