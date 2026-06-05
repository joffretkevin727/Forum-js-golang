package controller

import (
	"encoding/json"
	"forum/model"
	"net/http"
	"strconv"
)

type CommentController struct {
	Model *model.CommentModel
}

// HandleCommentsGET renvoie tous les commentaires pour que le front-end fasse le tri.
func (c *CommentController) HandleCommentsGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list, err := c.Model.GetAll()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(list)
}

// HandleCommentsPOST gère la création d'un commentaire.
func (c *CommentController) HandleCommentsPOST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body struct {
		Body     string `json:"body"`
		TopicID  int    `json:"topic_id"`
		AuthorID int    `json:"author_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}
	if err := c.Model.Create(body.Body, body.TopicID, body.AuthorID); err != nil {
		http.Error(w, "Erreur lors de la création", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// HandleCommentsDELETE gère la suppression d'un commentaire en exécutant directement la requête brute SQL.
func (c *CommentController) HandleCommentsDELETE(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	// Single line comment: executes a direct delete operation on the comments table by id.
	_, err = c.Model.DB.Exec(`DELETE FROM comments WHERE id = ?`, id)
	if err != nil {
		http.Error(w, "Erreur de suppression", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}