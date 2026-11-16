package main

import (
	"jul-spotify/handlers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env")
	}

	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Les variables d'environnement SPOTIFY_CLIENT_ID et SPOTIFY_CLIENT_SECRET doivent être définies.")
	}

	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("GET /{$}", handlers.HomeHandler)
	http.HandleFunc("/", handlers.NotFoundHandler)
	http.HandleFunc("GET /tracks/{id}", handlers.TracksHandler)
	log.Println("Starting server on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}