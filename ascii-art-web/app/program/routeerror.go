package asciiartweb

import "net/http"

type errorParams struct {
	Message string
}

func renderError(w http.ResponseWriter, msg string, code int) {
	// http.Error(w, msg, code)
	w.WriteHeader(code)
	err := &errorParams{Message: msg}
	renderTemplate(w, "error", *err)
}
