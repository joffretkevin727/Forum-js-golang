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

	// ============================================================
	// ROUTES : USER
	// ============================================================
	routeur.HandleFunc("POST /users/login", userController.LoginHandler)
	routeur.HandleFunc("POST /users", userController.CreateUserHandler)
	routeur.HandleFunc("GET /users/{username}", userController.GetUserHandler)
	routeur.HandleFunc("PUT /users/{id}/ban", userController.BanUserHandler)

	// ============================================================
	// ROUTES : TOPIC
	// ============================================================
	routeur.HandleFunc("GET /topics", topicController.GetTopicsHandler)
	routeur.HandleFunc("GET /topics/{id}", topicController.GetTopicHandler)
	routeur.HandleFunc("GET /setlike/{id}", topicController.SetLikeHandler)
	// Routes protégées par le Auth Middleware
	routeur.HandleFunc("POST /topics", AuthRequired(topicController.CreateTopicHandler))
	routeur.HandleFunc("PUT /topics/{id}", AuthRequired(topicController.UpdateTopicHandler))
	routeur.HandleFunc("DELETE /topics/{id}", AuthRequired(topicController.DeleteTopicHandler))

	// ============================================================
	// ROUTES : TAG
	// ============================================================
	routeur.HandleFunc("GET /tags", tagController.GetAllTagsHandler)
	routeur.HandleFunc("GET /topics/{topic_id}/tags", tagController.GetTagsByTopicHandler)
	// Routes protégées par le Auth Middleware
	routeur.HandleFunc("POST /tags", AuthRequired(tagController.CreateTagHandler))
	routeur.HandleFunc("POST /topics/tags", AuthRequired(tagController.AttachTagHandler))

	// ============================================================
	// ROUTES : MESSAGE
	// ============================================================
	routeur.HandleFunc("GET /topics/{topic_id}/messages", messageController.GetMessagesByTopicHandler)
	// Routes protégées par le Auth Middleware
	routeur.HandleFunc("POST /messages", AuthRequired(messageController.CreateMessageHandler))
	routeur.HandleFunc("DELETE /messages/{id}", AuthRequired(messageController.DeleteMessageHandler))

	// ============================================================
	// ROUTES : VOTE
	// ============================================================
	routeur.HandleFunc("GET /messages/{message_id}/score", voteController.GetScoreHandler)
	// Route protégée par le Auth Middleware
	routeur.HandleFunc("POST /votes", AuthRequired(voteController.VoteHandler))

	return routeur
}
