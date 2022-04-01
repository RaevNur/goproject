package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Universal parsing from api into data. Return 'true' if all okay
func parserAPI(api string, data interface{}) bool {
	// Get response
	resp, err := http.Get(api)
	if err != nil {
		log.Print(err.Error())
		return false
	}
	// Read body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err.Error())
		return false
	}
	// Unmarshal JSON
	err = json.Unmarshal(body, data)
	if err != nil {
		log.Print(err.Error())
		return false
	}
	return true
}

// Concert parser. Parse date to time.Time format and location coodinates
func parserConcert(concerts *[]Concert, datesLocations map[string][]string) {
	var tempConcert Concert
	for key, value := range datesLocations {
		address := key
		success := geocoding(address, &tempConcert)
		if !success {
			tempConcert.Location = key
			tempConcert.Coordinates[0] = 0.0
			tempConcert.Coordinates[1] = 0.0
		}
		tempConcert.Date = []time.Time{}
		tempConcert.DateFormatted = []string{}
		for _, timeValue := range value {
			date, err := time.Parse(timeLayout, timeValue)
			if err != nil {
				log.Print(err.Error())
			}
			tempConcert.Date = append(tempConcert.Date, date)
			tempConcert.DateFormatted = append(tempConcert.DateFormatted, timeValue)
		}
		*concerts = append(*concerts, tempConcert)
	}
}

// ShortArtists method
func (s *ShortArtists) Get(a *Artists) {
	for _, artist := range *a {
		*s = append(*s, ShortArtist{
			Id:    artist.Id,
			Name:  artist.Name,
			Image: artist.Image,
		})
	}
}
