package view

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

type ViewHandler struct {
	templates *template.Template
}

func newTemplate() (*template.Template, error) {
	// Gets All Templates in folder templates
	filepaths, err := filepath.Glob("templates/*.gohtml")
	files, err := template.ParseFiles(filepaths...)
	if err != nil {
		return nil, fmt.Errorf("newTemplate: %w", err)
	}
	return template.Must(files, nil), nil
}

func NewViewHandler() (*ViewHandler, error) {
	templates, err := newTemplate()
	if err != nil {
		return nil, fmt.Errorf("NewViewHandler: %w", err)
	}
	return &ViewHandler{
		templates: templates,
	}, nil
}

func (v *ViewHandler) InitRoutes(mux *http.ServeMux) {
	// HERE IS ALL ROUTES
	fsStatic := http.FileServer(http.Dir("templates/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fsStatic))

	// AnyRoutes
	mux.HandleFunc("/test", v.TestHandler)
}
