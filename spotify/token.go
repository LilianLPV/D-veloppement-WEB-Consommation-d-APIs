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

var httpClient = http.Client{
	Timeout: 5 * time.Second,
}

var tokenExpiration time.Time
var currentToken string

type SpotifyTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func RequestTokenSpotify() (SpotifyTokenResponse, error) {

	data := url.Values{}
	data.Add("client_id", os.Getenv("SPOTIFY_CLIENT_ID"))
	data.Add("client_secret", os.Getenv("SPOTIFY_CLIENT_SECRET"))
	data.Add("grant_type", "client_credentials")

	// Création de la requête POST avec les données encodées.
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return SpotifyTokenResponse{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		return SpotifyTokenResponse{}, err
	}

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

	// Retourne le token et nil
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
