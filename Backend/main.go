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
	// root:root correspond aux identifiants par défaut de MAMP
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/forum_db?parseTime=true")
	if err != nil {
		log.Fatal("Erreur de configuration DB:", err)
	}
	defer db.Close()

	// Vérifie si la connexion avec la base de données est réellement établie
	err = db.Ping()
	if err != nil {
		log.Fatal("Impossible de joindre la base de données:", err)
	}

	// ============================================================
	// INITIALISATION DU ROUTEUR
	// ============================================================
	apiRouter := router.New(db)

	// ============================================================
	// LANCEMENT DU SERVEUR HTTP
	// ============================================================
	log.Println("Serveur démarré sur http://localhost:6767")
	err = http.ListenAndServe(":6767", apiRouter)
	if err != nil {
		log.Fatal("Erreur lors du lancement du serveur:", err)
	}
}
