package spotify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// Le timeout va éviter que la requête bloque indéfiniment si Spotify ne répond pas.
var httpClient = http.Client{
	Timeout: 5 * time.Second,
}

var tokenExpiration time.Time
var currentToken string

// Les tags Json font le lien entre les clés du JSON et les champs de la structure.
type SpotifyTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

const APIBaseURL = "https://api.spotify.com/v1"

const maxRetries = 3

const maxRetryWait = 10 * time.Second

var (
	ErrNotFound    = errors.New("ressource introuvable sur Spotify")
	ErrUnavailable = errors.New("Spotify est indisponible")
)

// RequestTokenSpotify envoie une requête POST à l'API de Spotify pour obtenir un token d'accès.
func RequestTokenSpotify() (SpotifyTokenResponse, error) {

	//url.Values se charge de formater les données pour une requête POST.
	data := url.Values{}
	data.Add("client_id", os.Getenv("SPOTIFY_CLIENT_ID"))
	data.Add("client_secret", os.Getenv("SPOTIFY_CLIENT_SECRET"))
	data.Add("grant_type", "client_credentials")

	// Création de la requête POST avec les données encodées.
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return SpotifyTokenResponse{}, err
	}

	//Indique à Spotify que les données sont encodées en x-www-form-urlencoded et non Json.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	//Envoi de la requête et récupération de la réponse.
	resp, err := httpClient.Do(req)
	if err != nil {
		return SpotifyTokenResponse{}, err
	}

	// Ferme le body de la réponse à la fin de la fonction pour éviter les fuites de mémoire.
	defer resp.Body.Close()

	//Lecture complète du body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SpotifyTokenResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return SpotifyTokenResponse{}, fmt.Errorf("Erreur lors de la récupération du token: %s", resp.Status)
	}

	//Décodage du JSON dans la structure.
	var tokenResponse SpotifyTokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return SpotifyTokenResponse{}, err
	}

	// Retourne le token et nil pour indiquer qu'il n'y a pas d'erreur.
	return tokenResponse, nil
}

func GetToken() (string, error) {

	if currentToken != "" && time.Now().Add(time.Minute).Before(tokenExpiration) {
		return currentToken, nil
	}

	tokenResponse, err := RequestTokenSpotify()
	if err != nil {
		return "", err
	}

	tokenExpiration = time.Now().Add(time.Duration(tokenResponse.ExpiresIn) * time.Second)
	currentToken = tokenResponse.AccessToken

	return currentToken, nil
}

func Get(url string) ([]byte, error) {
	for attempt := 0; attempt <= maxRetries; attempt++ {
		token, err := GetToken()
		if err != nil {
			return nil, fmt.Errorf("%w: impossible d'obtenir un token: %v", ErrUnavailable, err)
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := httpClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrUnavailable, err)
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("%w: lecture de la réponse impossible: %v", ErrUnavailable, err)
		}

		switch {
		case resp.StatusCode == http.StatusOK:
			return body, nil

		case resp.StatusCode == http.StatusNotFound:
			return nil, ErrNotFound

		case resp.StatusCode == http.StatusTooManyRequests:
			wait := retryAfter(resp.Header.Get("Retry-After"))
			if wait > maxRetryWait {
				return nil, fmt.Errorf("%w: trop de requêtes, réessayer dans %s", ErrUnavailable, wait)
			}
			time.Sleep(wait)
			continue // je fais recommencer la requête

		case resp.StatusCode >= 500:
			return nil, fmt.Errorf("%w: %s", ErrUnavailable, resp.Status)

		default:
			return nil, fmt.Errorf("réponse inattendue de Spotify: %s - %s", resp.Status, string(body))
		}
	}

	return nil, fmt.Errorf("%w: toujours limité (429) après %d tentatives", ErrUnavailable, maxRetries)
}


// conversion de la valeur
func retryAfter(value string) time.Duration {
	seconds, err := strconv.Atoi(value)
	if err != nil || seconds < 0 {
		return time.Second
	}
	return time.Duration(seconds) * time.Second
}
