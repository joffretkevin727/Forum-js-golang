package router

import (
	"database/sql"
	"net/http"

	controller "forum/controller"
	model "forum/model"
)

// New configure toutes les routes de l'API et injecte les dépendances
func New(db *sql.DB) *http.ServeMux {
	//Initialisation des couches Model et Controller
	userModel := &model.UserModel{DB: db}
	userController := &controller.UserController{Model: userModel}

	topicModel := &model.TopicModel{DB: db}
	topicController := &controller.TopicController{Model: topicModel}

	//Création du multiplexeur (routeur)
	routeur := http.NewServeMux()

	//Déclaration des routes
	// USER
	routeur.HandleFunc("GET /user/{username}", userController.GetUserHandler)

	// TOPIC
	routeur.HandleFunc("GET /topic/{id}", topicController.GetTopicHandler)
	routeur.HandleFunc("GET /topics", topicController.GetTopicsHandler)
	// FAVORIS

	return routeur
}
