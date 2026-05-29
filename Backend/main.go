package main

import (
	"database/sql"
	"forum/router"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// ============================================================
	// CONNEXION À LA BASE DE DONNÉES
	// ============================================================
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum_db?parseTime=true&charset=utf8mb4")
	if err != nil {
		log.Fatal("Erreur de configuration DB:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Impossible de joindre la base de données:", err)
	}

	// ============================================================
	// INITIALISATION DU ROUTEUR ET MIDDLEWARES
	// ============================================================
	apiRouter := router.New(db)

	// Applique le middleware CORS global autour de ton routeur
	routerAvecCORS := router.EnableCORS(apiRouter)

	// ============================================================
	// LANCEMENT DU SERVEUR HTTP
	// ============================================================
	log.Println("Serveur démarré sur http://localhost:6767")
	// Utilise bien routerAvecCORS au lieu de apiRouter
	err = http.ListenAndServe(":6767", routerAvecCORS)
	if err != nil {
		log.Fatal("Erreur lors du lancement du serveur:", err)
	}
}
