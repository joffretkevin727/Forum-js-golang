package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "forum/model"
)

type TagController struct {
	Model *model.TagModel
}

// CreateTagHandler crée un tag unique
func (c *TagController) CreateTagHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}
	id, err := c.Model.Create(body.Name)
	if err != nil {
		http.Error(w, "Erreur ou tag déjà existant", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

// GetAllTagsHandler liste tous les tags
func (c *TagController) GetAllTagsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tags, err := c.Model.GetAll()
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tags)
}

// AttachTagHandler associe un tag existant à un topic
func (c *TagController) AttachTagHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		TopicID int `json:"topic_id"`
		TagID   int `json:"tag_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}
	if err := c.Model.AttachToTopic(body.TopicID, body.TagID); err != nil {
		http.Error(w, "Erreur d'association", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetTagsByTopicHandler liste les tags attachés à un sujet
func (c *TagController) GetTagsByTopicHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	topicID, err := strconv.Atoi(r.PathValue("topic_id"))
	if err != nil {
		http.Error(w, "ID de sujet invalide", http.StatusBadRequest)
		return
	}
	tags, err := c.Model.GetByTopic(topicID)
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tags)
}
