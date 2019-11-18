package controller

import "net/http"

// Home return OK
func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
