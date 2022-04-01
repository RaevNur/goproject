package api

import (
	"encoding/json"
	"log"
	"os"
)

// Get cache data from local file
func GetCache() bool {
	if AllArtists == nil {
		AllArtists = &Artists{}
	}
	file, err := os.Open(cachePath)
	if err != nil {
		return false
	}
	defer file.Close()
	temp := &Artists{}
	dec := json.NewDecoder(file)
	if err := dec.Decode(temp); err != nil {
		log.Print(err.Error())
		return false
	}
	AllArtists = temp
	log.Print("Data loaded from local cache")
	return true
}

// Get data from API and create local cache data
func CreateCache() bool {
	if !GetNewData() {
		log.Print("Can't create cache file without data")
		return false
	}
	*AllArtists = make(Artists, len(*NewAllArtists))
	copy(*AllArtists, *NewAllArtists)
	f, err := os.Create(cachePath)
	if err != nil {
		log.Print(err.Error())
		return false
	}
	jsonData, err := json.Marshal(AllArtists)
	if err != nil {
		log.Print(err.Error())
		f.Close()
		return false
	}
	f.Write(jsonData)
	log.Printf("Created '%v'", cachePath)
	f.Close()
	return true
}

// Rewrite local cache data
func rewriteCache(artists *Artists) bool {
	f, err := os.Create(cachePath)
	if err != nil {
		log.Print(err.Error())
		return false
	}
	jsonData, err := json.Marshal(artists)
	if err != nil {
		log.Print(err.Error())
		f.Close()
		return false
	}
	f.Write(jsonData)
	f.Close()
	return true
}
