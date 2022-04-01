package asciiartweb

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func initTemplates() {
	files, err := template.ParseFiles("app/templates/index.html", "app/templates/error.html")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	templates = template.Must(files, nil)
}

var templates *template.Template

func renderTemplate(w http.ResponseWriter, tmpl string, vars interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", vars)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}
