package router

import (
	"database/sql"
	"net/http"

	controller "forum/controller"
	model "forum/model"
)

// New configure toutes les routes de l'API et injecte les dépendances
func New(db *sql.DB) *http.ServeMux {
	// ============================================================
	// INITIALISATION DES MODÈLES
	// ============================================================
	userModel := &model.UserModel{DB: db}
	tagModel := &model.TagModel{DB: db}
	topicModel := &model.TopicModel{DB: db}
	messageModel := &model.MessageModel{DB: db}
	voteModel := &model.MessageVoteModel{DB: db}

	// ============================================================
	// INITIALISATION DES CONTRÔLEURS
	// ============================================================
	userController := &controller.UserController{Model: userModel}
	tagController := &controller.TagController{Model: tagModel}
	topicController := &controller.TopicController{Model: topicModel}
	messageController := &controller.MessageController{Model: messageModel}
	voteController := &controller.VoteController{Model: voteModel}

	// ============================================================
	// CONFIGURATION DU ROUTEUR
	// ============================================================
	routeur := http.NewServeMux()

	// GET    : Récupère des données depuis le serveur sans les modifier (ex: lire un sujet).
	// POST   : Envoie de nouvelles données au serveur pour créer une ressource (ex: publier un message).
	// PUT    : Remplace intégralement une ressource existante par les nouvelles données fournies.
	// DELETE : Supprime définitivement une ressource spécifique du serveur (ex: effacer un topic).

	// ============================================================
	// ROUTES : USER
	// ============================================================
	routeur.HandleFunc("POST /users", userController.CreateUserHandler)
	routeur.HandleFunc("GET /users/{username}", userController.GetUserHandler)
	routeur.HandleFunc("PUT /users/{id}/ban", userController.BanUserHandler)

	// ============================================================
	// ROUTES : TOPIC
	// ============================================================
	routeur.HandleFunc("POST /topics", topicController.CreateTopicHandler)
	routeur.HandleFunc("GET /topics", topicController.GetTopicsHandler)
	routeur.HandleFunc("GET /topics/{id}", topicController.GetTopicHandler)
	routeur.HandleFunc("PUT /topics/{id}", topicController.UpdateTopicHandler)
	routeur.HandleFunc("DELETE /topics/{id}", topicController.DeleteTopicHandler)

	// ============================================================
	// ROUTES : TAG
	// ============================================================
	routeur.HandleFunc("POST /tags", tagController.CreateTagHandler)
	routeur.HandleFunc("GET /tags", tagController.GetAllTagsHandler)
	routeur.HandleFunc("POST /topics/tags", tagController.AttachTagHandler)
	routeur.HandleFunc("GET /topics/{topic_id}/tags", tagController.GetTagsByTopicHandler)

	// ============================================================
	// ROUTES : MESSAGE
	// ============================================================
	routeur.HandleFunc("POST /messages", messageController.CreateMessageHandler)
	routeur.HandleFunc("GET /topics/{topic_id}/messages", messageController.GetMessagesByTopicHandler)
	routeur.HandleFunc("DELETE /messages/{id}", messageController.DeleteMessageHandler)

	// ============================================================
	// ROUTES : VOTE
	// ============================================================
	routeur.HandleFunc("POST /votes", voteController.VoteHandler)
	routeur.HandleFunc("GET /messages/{message_id}/score", voteController.GetScoreHandler)

	return routeur
}
