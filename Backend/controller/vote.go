package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "forum/model"
)

type VoteController struct {
	Model *model.MessageVoteModel
}

// VoteHandler enregistre ou change le vote d'un utilisateur sur un message
func (c *VoteController) VoteHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		UserID    int `json:"user_id"`
		MessageID int `json:"message_id"`
		Vote      int `json:"vote"` // Doit être +1 ou -1
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}
	if body.Vote != 1 && body.Vote != -1 {
		http.Error(w, "Valeur de vote invalide (doit être 1 ou -1)", http.StatusBadRequest)
		return
	}
	if err := c.Model.Vote(body.UserID, body.MessageID, body.Vote); err != nil {
		http.Error(w, "Erreur lors du vote", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetScoreHandler renvoie la somme cumulée des votes pour un message spécifique
func (c *VoteController) GetScoreHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	messageID, err := strconv.Atoi(r.PathValue("message_id"))
	if err != nil {
		http.Error(w, "ID de message invalide", http.StatusBadRequest)
		return
	}
	score, err := c.Model.GetScore(messageID)
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"score": score})
}
