package handlers

import (
	"html/template"
	"log"
	"net/http"

	"jul-spotify/spotify"
)

type trackCard struct {
	ID          string
	Name        string
	Cover       string
	ReleaseDate string
	Artists     string
	Duration    string
}

type homePage struct {
	Tracks  []trackCard
	Years   []string
	Filters Filters
	Total   int
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tracks, err := spotify.GetJulTracks()
	if err != nil {
		showError(w, err)
		return
	}

	filters := readFilters(r)
	filtered := filterTracks(tracks, filters)

	data := homePage{
		Years:   availableYears(tracks), 
		Filters: filters,
		Total:   len(filtered),
	}
	for _, track := range filtered {
		data.Tracks = append(data.Tracks, newTrackCard(track))
	}

	tmpl, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		log.Println("erreur de template:", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func newTrackCard(track spotify.Track) trackCard {
	card := trackCard{
		ID:          track.ID,
		Name:        track.Name,
		ReleaseDate: track.Album.ReleaseDate,
		Artists:     artistNames(track.Artists),
		Duration:    formatDuration(track.DurationMs),
	}
	images := track.Album.Images
	if len(images) > 1 {
		card.Cover = images[1].URL
	} else if len(images) == 1 {
		card.Cover = images[0].URL
	}
	return card
}