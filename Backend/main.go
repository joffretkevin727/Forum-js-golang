package main

import (
	"database/sql"
	"log"
	"net/http"

	"forum/router" // On importe notre nouveau package router

	_ "github.com/go-sql-driver/mysql" // Driver pour MAMP
)

func main() {
	//Connexion à la base de données MAMP
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/forum_db?parseTime=true")
	if err != nil {
		log.Fatal("Erreur de connexion DB:", err)
	}
	defer db.Close()

	//Initialisation du routeur en lui passant la DB
	apiRouter := router.New(db)

	//Lancement du serveur
	log.Println("Serveur démarré sur http://localhost:8080")
	err = http.ListenAndServe(":8080", apiRouter)
	if err != nil {
		log.Fatal("Erreur serveur:", err)
	}
}
