package controller

import (
	"encoding/json"
	"forum/model"
	"net/http"
	"strconv"
)

type TopicController struct {
	Model *model.TopicModel
}

// HandleTopicsGET récupère la totalité des sujets bruts stockés dans la base de données.
func (c *TopicController) HandleTopicsGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list, err := c.Model.GetAll()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(list)
}

// HandleTopicsPOST insère un nouveau sujet avec ses tags associés dans la base de données.
func (c *TopicController) HandleTopicsPOST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body struct {
		Title    string   `json:"title"`
		Body     string   `json:"text"`
		Pseudo   string   `json:"pseudo"`
		AuthorID int      `json:"author_id"`
		Tags     []string `json:"tags"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}
	// Single line comment: calls model create method with parsed parameters.
	id, err := c.Model.Create(body.Title, body.Body, body.Pseudo, body.AuthorID, body.Tags)
	if err != nil {
		http.Error(w, "Erreur lors de la création", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

// HandleTopicsPUT modifie les champs d'un sujet existant ciblé par son identifiant.
func (c *TopicController) HandleTopicsPUT(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	var body struct {
		Title  string `json:"title"`
		Body   string `json:"text"`
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}
	// Single line comment: updates topic fields directly using SQL database exec.
	_, err = c.Model.DB.Exec("UPDATE topics SET title = ?, body = ?, status = ? WHERE id = ?", body.Title, body.Body, body.Status, id)
	if err != nil {
		http.Error(w, "Erreur de mise à jour", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// HandleTopicsDELETE supprime définitivement de la base de données le sujet correspondant à l'identifiant.
func (c *TopicController) HandleTopicsDELETE(w http.ResponseWriter, r *http.Request) {
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

func (c *TopicController) HandleTopicIDGET(w http.ResponseWriter, r *http.Request)    {}
func (c *TopicController) HandleTopicIDPOST(w http.ResponseWriter, r *http.Request)   {}
func (c *TopicController) HandleTopicIDPUT(w http.ResponseWriter, r *http.Request)    {}
func (c *TopicController) HandleTopicIDDELETE(w http.ResponseWriter, r *http.Request) {}
