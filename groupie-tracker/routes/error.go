package routes

import "net/http"

func renderError(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	renderTemplate(w, "error", http.StatusText(code))
}
