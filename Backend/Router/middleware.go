package router

import "net/http"

// EnableCORS ajoute les entêtes nécessaires pour autoriser le Front à requêter l'API
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-User-ID")

		// Gère la requête de pré-vérification (Preflight)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// AuthRequired vérifie si la requête contient un identifiant utilisateur valide
func AuthRequired(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Vérification simple via un Header (adaptable plus tard avec un Token)
		userID := r.Header.Get("X-User-ID")
		if userID == "" {
			http.Error(w, `{"message":"Accès refusé : vous devez être connecté"}`, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
