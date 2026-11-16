package spotify

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type ExternalURLs struct {
	Spotify string `json:"spotify"`
}

type Artist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Album struct {
	ID                   string       `json:"id"`
	Name                 string       `json:"name"`
	AlbumType            string       `json:"album_type"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	Images               []Image      `json:"images"`
	Artists              []Artist     `json:"artists"`
	ExternalURLs         ExternalURLs `json:"external_urls"`
}

type Track struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	DurationMs   int          `json:"duration_ms"`
	Artists      []Artist     `json:"artists"`
	Album        Album        `json:"album"`
	ExternalURLs ExternalURLs `json:"external_urls"`
}

type AlbumsPage struct {
	Items []Album `json:"items"`
	Next  string  `json:"next"`
}

type TracksPage struct {
	Items []Track `json:"items"`
	Next  string  `json:"next"`
}