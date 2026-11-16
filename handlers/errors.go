package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// errorPage contient ce qu'affiche la page d'erreur : le code et un message.
type errorPage struct {
	Code    int
	Title   string
	Message string
}

// affiche la page 404 pour toutes les routes inconnues
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	showNotFound(w)
}

func showNotFound(w http.ResponseWriter) {
	showErrorPage(w, errorPage{
		Code:    http.StatusNotFound,
		Title:   "Page introuvable",
		Message: "Cette page ou ce titre n'existe pas.",
	})
}

func showError(w http.ResponseWriter, spotifyErr error) {
	log.Println("erreur Spotify:", spotifyErr)
	showErrorPage(w, errorPage{
		Code:    http.StatusBadGateway,
		Title:   "Oups…",
		Message: "Spotify ne répond pas pour le moment. Réessaie dans quelques instants.",
	})
}


func showErrorPage(w http.ResponseWriter, page errorPage) {
	tmpl, err := template.ParseFiles("templates/error.gohtml")
	if err != nil {
		log.Println("erreur de template:", err)
		http.Error(w, page.Message, page.Code)
		return
	}
	w.WriteHeader(page.Code)
	tmpl.Execute(w, page)
}