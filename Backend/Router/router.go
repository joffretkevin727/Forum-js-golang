package router

import (
	"database/sql"
	"net/http"

	"forum/controller"
	"forum/model"
)

func New(db *sql.DB) *http.ServeMux {
	userModel := &model.UserModel{DB: db}
	tagModel := &model.TagModel{DB: db}
	topicModel := &model.TopicModel{DB: db}
	commentModel := &model.CommentModel{DB: db}
	likeModel := &model.LikeModel{DB: db}

	userController := &controller.UserController{Model: userModel}
	tagController := &controller.TagController{Model: tagModel}
	topicController := &controller.TopicController{Model: topicModel}
	commentController := &controller.CommentController{Model: commentModel}
	likeController := &controller.LikeController{Model: likeModel}

	routeur := http.NewServeMux()

	routeur.HandleFunc("POST /users/login", userController.LoginHandler)
	routeur.HandleFunc("POST /users", userController.CreateUserHandler)
	routeur.HandleFunc("PUT /users/ban", userController.BanUserHandler)
	routeur.HandleFunc("DELETE /users", userController.DeleteUserHandler)

	routeur.HandleFunc("GET /topics", topicController.HandleTopicsGET)
	routeur.HandleFunc("POST /topics", topicController.HandleTopicsPOST)
	routeur.HandleFunc("PUT /topics", topicController.HandleTopicsPUT)
	routeur.HandleFunc("DELETE /topics", topicController.HandleTopicsDELETE)

	routeur.HandleFunc("POST /topics/like", likeController.VoteTopicHandler)

	routeur.HandleFunc("GET /comments", commentController.HandleCommentsGET)
	routeur.HandleFunc("POST /comments", commentController.HandleCommentsPOST)
	routeur.HandleFunc("DELETE /comments", commentController.HandleCommentsDELETE)
	routeur.HandleFunc("POST /comments/like", likeController.VoteCommentHandler)
	routeur.HandleFunc("GET /topiccomments", commentController.HandleCommentsByIdGET)

	routeur.HandleFunc("GET /tags", tagController.GetAllTagsHandler)
	routeur.HandleFunc("POST /tags", tagController.CreateTagHandler)

	return routeur
}
