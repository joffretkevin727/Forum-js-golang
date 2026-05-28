package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "forum/model"
)

type MessageController struct {
	Model *model.MessageModel
}

// CreateMessageHandler publie un nouveau commentaire/message
func (c *MessageController) CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
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
	id, err := c.Model.Create(body.Body, body.TopicID, body.AuthorID)
	if err != nil {
		http.Error(w, "Erreur lors de la création", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

// GetMessagesByTopicHandler récupère le fil de discussion complet d'un sujet
func (c *MessageController) GetMessagesByTopicHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	topicID, err := strconv.Atoi(r.PathValue("topic_id"))
	if err != nil {
		http.Error(w, "ID de sujet invalide", http.StatusBadRequest)
		return
	}
	messages, err := c.Model.GetByTopic(topicID)
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(messages)
}

// DeleteMessageHandler supprime un commentaire
func (c *MessageController) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
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
