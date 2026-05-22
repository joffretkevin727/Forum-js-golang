package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"forum/model"
)

type TopicController struct {
	Model *model.TopicModel
}

// GetTopicHandler récupère un seul sujet par son ID
func (controleur *TopicController) GetTopicHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "Le paramètre ID est requis", http.StatusBadRequest)
		return
	}

	// CONVERSION : Convertit la chaîne en entier
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "L'ID doit être un nombre valide", http.StatusBadRequest)
		return
	}

	// Appel au Modèle (on passe la variable id propre)
	topic, err := controleur.Model.GetByID(id)
	if err != nil {
		log.Printf("Erreur lors de la récupération du topic %d: %v", id, err)
		http.Error(w, "Topic introuvable", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(topic)
}

// GetTopicsHandler récupère une liste de sujets paginée
func (controleur *TopicController) GetTopicsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. Récupération des paramètres de l'URL (?page=X&limit=Y)
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	// 2. Valeurs par défaut si l'utilisateur ne précise rien
	page := 1
	limit := 10

	// 3. Conversion de la page en entier
	if pageStr != "" {
		pageValide, err := strconv.Atoi(pageStr)
		if err == nil && pageValide > 0 {
			page = pageValide
		}
	}

	// 4. Conversion de la limite en entier
	if limitStr != "" {
		limiteValide, err := strconv.Atoi(limitStr)
		if err == nil && limiteValide > 0 {
			limit = limiteValide
		}
	}

	// 5. Calcul du décalage (l'offset) pour MySQL
	offset := (page - 1) * limit

	// 6. Appel au modèle avec des paramètres clairs
	listeTopics, err := controleur.Model.GetMany(limit, offset)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération", http.StatusInternalServerError)
		return
	}

	// 7. Envoi de la liste en JSON
	json.NewEncoder(w).Encode(listeTopics)
}
