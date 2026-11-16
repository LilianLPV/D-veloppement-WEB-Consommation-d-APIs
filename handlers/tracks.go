package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"jul-spotify/spotify"
)

type trackPage struct {
	Name        string
	AlbumName   string
	Artists     string
	ReleaseDate string
	Duration    string
	Cover       string
	SpotifyURL  string
}

func TracksHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	track, err := spotify.GetTrack(id)
	if errors.Is(err, spotify.ErrNotFound) {
		showNotFound(w)
		return
	}
	if err != nil {
		showError(w, err)
		return
	}

	data := trackPage{
		Name:        track.Name,
		AlbumName:   track.Album.Name,
		Artists:     artistNames(track.Artists),
		ReleaseDate: track.Album.ReleaseDate,
		Duration:    formatDuration(track.DurationMs),
		SpotifyURL:  track.ExternalURLs.Spotify,
	}
	if len(track.Album.Images) > 0 {
		data.Cover = track.Album.Images[0].URL
	}

	tmpl, err := template.ParseFiles("templates/tracks.gohtml")
	if err != nil {
		log.Println("erreur de template:", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

// Conversion de la durée des titres
func formatDuration(ms int) string {
	totalSeconds := ms / 1000
	return fmt.Sprintf("%02d:%02d", totalSeconds/60, totalSeconds%60)
}

// affichage des noms d'artistes
func artistNames(artists []spotify.Artist) string {
	names := make([]string, 0, len(artists))
	for _, a := range artists {
		names = append(names, a.Name)
	}
	return strings.Join(names, ", ")
}