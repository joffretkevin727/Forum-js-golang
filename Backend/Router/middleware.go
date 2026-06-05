package router

import (
        "net/http"
        "strconv"
        "context"
        )

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
// Single line comment: Denies access with an explicit JSON message if the X-User-ID header is missing.
func AuthRequired(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
       userIDStr := r.Header.Get("X-User-ID")
       if userIDStr == "" {
          w.Header().Set("Content-Type", "application/json")
          w.WriteHeader(http.StatusUnauthorized)
          w.Write([]byte(`{"message":"Accès refusé : vous devez être connecté avec un compte valide pour créer un sujet"}`))
          return
       }

       userID, err := strconv.Atoi(userIDStr)
       if err != nil {
          w.Header().Set("Content-Type", "application/json")
          w.WriteHeader(http.StatusBadRequest)
          w.Write([]byte(`{"message":"ID utilisateur invalide"}`))
          return
       }

       ctx := context.WithValue(r.Context(), "userID", userID)
       next.ServeHTTP(w, r.WithContext(ctx))
    }
}
