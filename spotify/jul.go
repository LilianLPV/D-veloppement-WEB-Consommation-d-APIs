package spotify

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

const JulArtistID = "3IW7ScrzXmPvZhB27hmfgy"

const cacheDuration = 24 * time.Hour

const cacheFile = "cache/jul_tracks.json"

// Cache en mémoire des titres du jul
var (
	cachedTracks    []Track
	cacheExpiration time.Time
)

// ma strucutre pour le cache
type diskCache struct {
	SavedAt time.Time `json:"saved_at"`
	Tracks  []Track   `json:"tracks"`
}

func GetTrack(id string) (Track, error) {

	if !isValidID(id) {
		return Track{}, ErrNotFound
	}

	body, err := Get(APIBaseURL + "/tracks/" + id)
	if err != nil {
		return Track{}, err
	}

	var track Track
	if err := json.Unmarshal(body, &track); err != nil {
		return Track{}, fmt.Errorf("%w: réponse illisible: %v", ErrUnavailable, err)
	}
	return track, nil
}

func GetJulTracks() ([]Track, error) {

	if cachedTracks != nil && time.Now().Before(cacheExpiration) {
		return cachedTracks, nil
	}

	if tracks, savedAt, ok := loadDiskCache(); ok && time.Since(savedAt) < cacheDuration {
		cachedTracks = tracks
		cacheExpiration = savedAt.Add(cacheDuration)
		return tracks, nil
	}

	log.Println("récupération des titres de JUL auprès de Spotify…")
	tracks, err := fetchJulTracks()
	if err != nil {
		return nil, err
	}

	cachedTracks = tracks
	cacheExpiration = time.Now().Add(cacheDuration)
	saveDiskCache(tracks)
	log.Println(len(tracks), "titres récupérés et mis en cache")
	return tracks, nil
}

func fetchJulTracks() ([]Track, error) {
	albums, err := getJulAlbums()
	if err != nil {
		return nil, err
	}

	var tracks []Track
	for _, album := range albums {
		albumTracks, err := getAlbumTracks(album.ID)
		if err != nil {
			return nil, err
		}

		for _, track := range albumTracks {
			track.Album = album
			tracks = append(tracks, track)
		}
	}

	sort.SliceStable(tracks, func(i, j int) bool {
		return tracks[i].Album.ReleaseDate > tracks[j].Album.ReleaseDate
	})
	return tracks, nil
}

func loadDiskCache() (tracks []Track, savedAt time.Time, ok bool) {
	data, err := os.ReadFile(cacheFile)
	if err != nil {
		return nil, time.Time{}, false
	}
	var c diskCache
	if err := json.Unmarshal(data, &c); err != nil || len(c.Tracks) == 0 {
		return nil, time.Time{}, false
	}
	return c.Tracks, c.SavedAt, true
}

func saveDiskCache(tracks []Track) {
	data, err := json.Marshal(diskCache{SavedAt: time.Now(), Tracks: tracks})
	if err == nil {
		err = os.MkdirAll("cache", 0o755)
	}
	if err == nil {
		err = os.WriteFile(cacheFile, data, 0o644)
	}
	if err != nil {
		log.Println("impossible d'écrire le cache:", err)
	}
}

func getJulAlbums() ([]Album, error) {
	var albums []Album
	url := APIBaseURL + "/artists/" + JulArtistID + "/albums?include_groups=album,single"

	for url != "" {
		body, err := Get(url)
		if err != nil {
			return nil, err
		}

		var page AlbumsPage
		if err := json.Unmarshal(body, &page); err != nil {
			return nil, fmt.Errorf("%w: réponse illisible: %v", ErrUnavailable, err)
		}

		albums = append(albums, page.Items...)
		url = page.Next // "" quand il n'y a plus de page
	}
	return albums, nil
}

func getAlbumTracks(albumID string) ([]Track, error) {
	var tracks []Track
	url := APIBaseURL + "/albums/" + albumID + "/tracks?limit=50"

	for url != "" {
		body, err := Get(url)
		if err != nil {
			return nil, err
		}

		var page TracksPage
		if err := json.Unmarshal(body, &page); err != nil {
			return nil, fmt.Errorf("%w: réponse illisible: %v", ErrUnavailable, err)
		}

		tracks = append(tracks, page.Items...)
		url = page.Next
	}
	return tracks, nil
}

func isValidID(id string) bool {
	if len(id) != 22 {
		return false
	}
	for _, c := range id {
		isLetter := (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
		isDigit := c >= '0' && c <= '9'
		if !isLetter && !isDigit {
			return false
		}
	}
	return true
}
