package asciiartweb

import (
	"asciiartweb/app/asciiart"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type PageVars struct {
	AvailableFonts map[string]bool
	Font           string
	Input          string
	Result         string
}

func (pv *PageVars) Constructor(r *http.Request) {
	pv.Font = r.FormValue("font")
	pv.Input = formatForASCII(r.FormValue("input"))
	pv.AvailableFonts = Fonts
}

func (pv *PageVars) getASCIIConfigs() *asciiart.ASCIIConfigs {
	return &asciiart.ASCIIConfigs{FontName: pv.Font, Text: pv.Input}
}

// ASCIIArtHandler - for path "/ascii-art", takes only post method
func ASCIIArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		renderError(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// Check method
	if r.Method != http.MethodPost {
		renderError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	// Parse values
	err := r.ParseForm()
	if err != nil {
		renderError(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	pageVars := &PageVars{}
	pageVars.Constructor(r)
	// Check font
	if _, isIn := Fonts[pageVars.Font]; !isIn {
		renderError(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// Check button clicked
	submit := r.FormValue("submit")
	if submit != "show" && submit != "download" {
		renderError(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// Check input
	if !asciiart.IsValidInputText(pageVars.Input) {
		renderError(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// Generate ascii art
	asciiConfigs := pageVars.getASCIIConfigs()
	pageVars.Result, err = asciiart.GetArtByConfigs(asciiConfigs)
	if err != nil {
		renderError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	switch submit {
	case "show":
		renderTemplate(w, "index", *pageVars)
	case "download":
		file := strings.NewReader(pageVars.Result)
		fileSize := strconv.FormatInt(file.Size(), 10)
		w.Header().Set("Content-Disposition", "attachment; filename=asciiart.txt")
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", fileSize)
		file.Seek(0, 0)
		io.Copy(w, file)
	}
}

func formatForASCII(str string) string {
	return strings.ReplaceAll(str, "\r", "")
}
