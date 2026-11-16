package handlers

import (
	"net/http"

	"jul-spotify/spotify"
)

type Filters struct {
	Year string
	Type string
	Feat string
}

func readFilters(r *http.Request) Filters {
	query := r.URL.Query()
	f := Filters{Year: query.Get("year")}

	if t := query.Get("type"); t == "single" || t == "album" {
		f.Type = t
	}
	if feat := query.Get("feat"); feat == "oui" || feat == "non" {
		f.Feat = feat
	}
	return f
}

func filterTracks(tracks []spotify.Track, f Filters) []spotify.Track {
	var result []spotify.Track

	for _, track := range tracks {
		if f.Year != "" && releaseYear(track) != f.Year {
			continue
		}
		if f.Type != "" && track.Album.AlbumType != f.Type {
			continue
		}
		if f.Feat == "oui" && !hasFeat(track) {
			continue
		}
		if f.Feat == "non" && hasFeat(track) {
			continue
		}
		result = append(result, track)
	}
	return result
}

func hasFeat(track spotify.Track) bool {
	for _, artist := range track.Artists {
		if artist.ID != spotify.JulArtistID {
			return true
		}
	}
	return false
}

func releaseYear(track spotify.Track) string {
	date := track.Album.ReleaseDate
	if len(date) < 4 {
		return ""
	}
	return date[:4]
}

func availableYears(tracks []spotify.Track) []string {
	var years []string
	seen := map[string]bool{}

	for _, track := range tracks {
		year := releaseYear(track)
		if year != "" && !seen[year] {
			seen[year] = true
			years = append(years, year)
		}
	}
	return years
}
