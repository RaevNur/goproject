package view

import (
	"fmt"
	"log"
	"net/http"
)

// TestHandler - Handle for Testing
func (v *ViewHandler) TestHandler(w http.ResponseWriter, r *http.Request) {
	err := v.debugRefreshTemplates()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// err = v.templates.ExecuteTemplate(w, "pg-index", nil)
	// err = v.templates.ExecuteTemplate(w, "pg-signup", nil)
	// err = v.templates.ExecuteTemplate(w, "pg-login", nil)
	// err = v.templates.ExecuteTemplate(w, "pg-question", nil)
	// err = v.templates.ExecuteTemplate(w, "pg-questions", nil)
	// err = v.templates.ExecuteTemplate(w, "pg-question-create", nil)
	// err = v.templates.ExecuteTemplate(w, "pg-tags", nil)
	// err = v.templates.ExecuteTemplate(w, "pg-user", nil)
	err = v.templates.ExecuteTemplate(w, "pg-users", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//? debugRefreshTemplates -
func (v *ViewHandler) debugRefreshTemplates() error {
	templates, err := newTemplate()
	if err != nil {
		return fmt.Errorf("debugRefreshTemplates: %w", err)
	}
	v.templates = templates
	return nil
}
