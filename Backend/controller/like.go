package controller

import (
	"encoding/json"
	"forum/model"
	"net/http"
)

type LikeController struct {
	Model *model.LikeModel
}

// VoteTopicHandler enregistre ou met à jour le vote d'un utilisateur sur un sujet spécifique.
func (c *LikeController) VoteTopicHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		UserID  int `json:"user_id"`
		TopicID int `json:"topic_id"`
		Like    int `json:"like"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}
	// Single line comment: triggers generic model method for topic votes.
	if err := c.Model.VoteTarget("liketopic", "topicid", body.UserID, body.TopicID, body.Like); err != nil {
		http.Error(w, "Erreur lors du vote", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// VoteCommentHandler enregistre ou met à jour le vote d'un utilisateur sur un commentaire spécifique.
func (c *LikeController) VoteCommentHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		UserID    int `json:"user_id"`
		CommentID int `json:"comment_id"`
		Like      int `json:"like"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}
	// Single line comment: triggers generic model method for comment votes.
	if err := c.Model.VoteTarget("likecomment", "commentid", body.UserID, body.CommentID, body.Like); err != nil {
		http.Error(w, "Erreur lors du vote", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}