package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
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
