package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "forum/model"
)

type TopicController struct {
	Model *model.TopicModel
}

// GetTopicHandler récupère un sujet par son ID
func (c *TopicController) GetTopicHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	topic, err := c.Model.GetByID(id)
	if err != nil {
		http.Error(w, "Sujet introuvable", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(topic)
}

// GetTopicsHandler récupère la liste des sujets avec pagination
func (c *TopicController) GetTopicsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	limite, _ := strconv.Atoi(r.URL.Query().Get("limite"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	if limite <= 0 {
		limite = 10
	}
	topics, err := c.Model.GetMany(limite, offset)
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(topics)
}

// CreateTopicHandler insère un nouveau sujet
// Single line comment: Decodes JSON payload including tags and returns the new auto-incremented ID.
func (c *TopicController) CreateTopicHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var body struct {
       Title    string   `json:"title"`
       Body     string   `json:"body"`
       AuthorID int      `json:"author_id"`
       Tags     []string `json:"tags"`
    }
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
       http.Error(w, "Format JSON invalide", http.StatusBadRequest)
       return
    }
    id, err := c.Model.Create(body.Title, body.Body, body.AuthorID, body.Tags)
    if err != nil {
       http.Error(w, "Erreur lors de la création", http.StatusInternalServerError)
       return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

// UpdateTopicHandler modifie un sujet existant
func (c *TopicController) UpdateTopicHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	var body struct {
		Title  string `json:"title"`
		Body   string `json:"body"`
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}
	if err := c.Model.Update(id, body.Title, body.Body, body.Status); err != nil {
		http.Error(w, "Erreur de mise à jour", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DeleteTopicHandler supprime un sujet
func (c *TopicController) DeleteTopicHandler(w http.ResponseWriter, r *http.Request) {
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
