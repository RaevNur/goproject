package asciiartweb

import (
	"net/http"
)

// IndexHandler - its for Index page "/"
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// Check method
	if r.Method != http.MethodGet {
		renderError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	renderTemplate(w, "index", PageVars{Font: "standard", AvailableFonts: Fonts})
}
