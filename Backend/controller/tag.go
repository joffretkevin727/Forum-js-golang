package controller

import (
	"encoding/json"
	"forum/model"
	"net/http"
)

type TagController struct {
	Model *model.TagModel
}

// GetAllTagsHandler récupère tous les tags disponibles pour alimenter l'autocomplétion du formulaire front-end.
func (c *TagController) GetAllTagsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tags, err := c.Model.GetAll()
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tags)
}

// CreateTagHandler enregistre un nouveau tag unique créé par l'utilisateur en base de données.
func (c *TagController) CreateTagHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}
	// Single line comment: executes a direct insert operation into the tags table using DB connection.
	_, err := c.Model.DB.Exec(`INSERT INTO tags (name) VALUES (?)`, body.Name)
	if err != nil {
		http.Error(w, "Erreur ou tag déjà existant", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}