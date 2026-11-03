package handlers

import "net/http"

func TracksHandler(w http.ResponseWriter, r *http.Request) {
	
	w.Write([]byte("Tracks endpoint"))
}
