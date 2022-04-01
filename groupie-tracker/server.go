package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"groupie-tracker/api"
	"groupie-tracker/routes"
)

func getData(ch chan bool, duration time.Duration) {
	t := time.Now()
	isCacheLoaded := api.GetCache()
	isDataUpdated := false
	if !isCacheLoaded {
		isDataUpdated = api.CreateCache()
		if !isDataUpdated {
			ch <- false
			close(ch)
			return
		}
	}
	api.GetSuggestions()
	api.GetCities()
	log.Printf("Data parsed in %v", time.Since(t))
	ch <- true
	close(ch)

	if isDataUpdated {
		time.Sleep(duration)
	}
	for {
		log.Print("Start parse new Data")
		t = time.Now()
		api.RefreshAll()
		log.Printf("New data parsed in %v", time.Since(t))
		time.Sleep(duration)
	}
}

func runServer() {
	routes.InitTemplates()
	mux := http.NewServeMux()
	// Fs Assets
	assets := http.FileServer(http.Dir("public/assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", assets))
	// Front Pages
	mux.HandleFunc("/", routes.IndexHandler) // Index Page
	// API Pages
	mux.HandleFunc("/api", routes.ApiHandler)                     // How Use API
	mux.HandleFunc("/api/artists", routes.ArtistsHandler)         // All Artists
	mux.HandleFunc("/api/artists/", routes.ArtistHandler)         // Artist By Id
	mux.HandleFunc("/api/filter", routes.FilterHandler)           // Api Filter
	mux.HandleFunc("/api/suggestions", routes.SuggestionsHandler) // Api Suggestions
	mux.HandleFunc("/api/cities", routes.CitiesHandler)           // Api Cities
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	addr := fmt.Sprintf("localhost:%v", port)
	// Start Listen
	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Printf("Server started on http://%v", addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Print(err.Error())
	}
}
