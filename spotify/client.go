package spotify

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const APIBaseURL = "https://api.spotify.com/v1"

const maxRetries = 3

const maxRetryWait = 10 * time.Second

var (
	ErrNotFound    = errors.New("ressource introuvable sur Spotify")
	ErrUnavailable = errors.New("Spotify est indisponible")
)

var blockedUntil time.Time

const minDelay = 250 * time.Millisecond

var lastRequest time.Time

func Get(url string) ([]byte, error) {
	if time.Now().Before(blockedUntil) {
		return nil, fmt.Errorf("%w: limité par Spotify jusqu'à %s", ErrUnavailable, blockedUntil.Format("02/01 15:04"))
	}

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

		if wait := minDelay - time.Since(lastRequest); wait > 0 {
			time.Sleep(wait)
		}
		lastRequest = time.Now()

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
				blockedUntil = time.Now().Add(wait)
				return nil, fmt.Errorf("%w: trop de requêtes, réessayer dans %s", ErrUnavailable, wait)
			}
			time.Sleep(wait)
			continue // recommence la requête

		case resp.StatusCode >= 500:
			return nil, fmt.Errorf("%w: %s", ErrUnavailable, resp.Status)

		default:
			return nil, fmt.Errorf("réponse inattendue de Spotify: %s - %s", resp.Status, string(body))
		}
	}

	return nil, fmt.Errorf("%w: toujours limité (429) après %d tentatives", ErrUnavailable, maxRetries)
}

func retryAfter(value string) time.Duration {
	seconds, err := strconv.Atoi(value)
	if err != nil || seconds < 0 {
		return time.Second
	}
	return time.Duration(seconds) * time.Second
}