package handlers

import "net/http"

const message = "Hello World!"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
