package main

import (
	"database/sql"
	"fmt"
	"forum/router"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv" // Single line comment: imported for loading .env variables
)

func main() {
	// Charge le fichier .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env")
	}

	// Récupère les variables d'environnement
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	apiPort := os.Getenv("PORT")

	// Construit la chaîne de connexion dynamiquement
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Connexion à la base de données
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Erreur de configuration DB:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Impossible de joindre la base de données:", err)
	}

	apiRouter := router.New(db)
	routerAvecCORS := router.EnableCORS(apiRouter)

	log.Printf("=====================================================================================================================")
	log.Printf("base de donnée accessible ici: http://localhost/phpMyAdmin5/index.php?route=/database/structure&server=1&db=forum_db ")
	log.Printf("=====================================================================================================================")
	log.Printf("")
	log.Printf("")
	log.Printf("==================================================")
	log.Printf("Serveur démarré sur http://localhost:%s", apiPort)
	log.Printf("==================================================")
	err = http.ListenAndServe(":"+apiPort, routerAvecCORS)
	if err != nil {
		log.Fatal("Erreur lors du lancement du serveur:", err)
	}
}
