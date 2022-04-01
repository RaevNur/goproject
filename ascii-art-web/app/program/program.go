package asciiartweb

import (
	"log"
	"net/http"
	"time"
)

var gConfigs *gconfigs = &gconfigs{HostURL: ":8080"}

func init() {
	SetGlobalConfigs()
	InitFonts()
	initTemplates()
}

// Program - Execute Program With Default Json Configs
func Program() {
	ListenRoutes()
}

// ListenRoutes - Starting Listening All Routes in routes
func ListenRoutes() {
	log.Printf("Start Listening Server: http://localhost%v", gConfigs.HostURL)
	mux := http.NewServeMux()
	//Route For CSS
	fileServer := http.FileServer(http.Dir("app/templates/css/"))
	mux.Handle("/css/", http.StripPrefix("/css/", fileServer))
	// Add Routes
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/ascii-art", ASCIIArtHandler)
	// Start Listen
	server := http.Server{
		Addr:         gConfigs.HostURL,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf(err.Error())
	}
}
