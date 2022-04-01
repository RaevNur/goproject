package api

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

// Refresh data in AllArtists
func RefreshAll() {
	if !GetNewData() {
		log.Print("Can't refresh Artists cause something wrong with API")
		return
	}
	rewriteCache(NewAllArtists)
	*AllArtists = make(Artists, len(*NewAllArtists))
	copy(*AllArtists, *NewAllArtists)
	GetSuggestions()
	GetCities()
}

// Get new data from API
func GetNewData() bool {
	// Flush NewAllArtists
	NewAllArtists = &Artists{}
	// At first parse artists data
	var artists []artistsJSON
	var relations relationsJSON
	success := parserAPI(artistsAPI, &artists)
	if !success {
		return success
	}
	success = parserAPI(relationAPI, &relations)
	if !success {
		return success
	}
	// For every artist get relation info and append it to All
	for ind, artist := range artists {
		tempArtist := Artist{
			Id:                  artist.Id,
			Name:                artist.Name,
			Image:               artist.Image,
			Members:             artist.Members,
			CreationDate:        artist.CreationDate,
			FirstAlbumFormatted: artist.FirstAlbum,
			Concerts:            []Concert{},
		}
		date, err := time.Parse(timeLayout, artist.FirstAlbum)
		if err != nil {
			log.Print(err.Error())
		}
		tempArtist.FirstAlbum = date
		// Parse relation
		// Maybe change after (to do logic)
		// if artist.Id == relations.Index[ind].Id {
		parserConcert(&tempArtist.Concerts, relations.Index[ind].DatesLocations)
		// }
		*NewAllArtists = append(*NewAllArtists, tempArtist)
	}
	return true
}

// Get suggestions for autofill or suggestion
func GetSuggestions() {
	AllSuggestions = &Suggestions{}
	temp := make(map[Suggestion]bool)
	for _, artist := range *AllArtists {
		temp[Suggestion{Name: artist.Name, Type: "Name"}] = true
		for _, member := range artist.Members {
			temp[Suggestion{Name: member, Type: "Member"}] = true
		}
		temp[Suggestion{Name: strconv.Itoa(artist.CreationDate), Type: "Creation Date"}] = true
		temp[Suggestion{Name: artist.FirstAlbumFormatted, Type: "First Album Date"}] = true
		for _, concerts := range artist.Concerts {
			temp[Suggestion{Name: concerts.Location, Type: "Concert Location"}] = true
		}
	}
	for key := range temp {
		*AllSuggestions = append(*AllSuggestions, key)
	}
	sort.Slice(*AllSuggestions, func(i, j int) bool {
		return (*AllSuggestions)[i].Name < (*AllSuggestions)[j].Name
	})
}

// Get cities
func GetCities() {
	AllCities = &[]string{}
	temp := make(map[string]bool)
	for _, artist := range *AllArtists {
		for _, concert := range artist.Concerts {
			temp[concert.Location] = true
		}
	}
	for key := range temp {
		*AllCities = append(*AllCities, key)
	}
	sort.Slice(*AllCities, func(i, j int) bool {
		return (*AllCities)[i] < (*AllCities)[j]
	})
}

// MAP API : Gives formatted location name and coordinates by the city name
func geocoding(city string, concert *Concert) bool {
	url := fmt.Sprintf(geocodingURL, city, googleMapAPI)
	var data geocodingJSON
	success := parserAPI(url, &data)
	if !success || len(data.Results) < 1 {
		return false
	}
	concert.Location = data.Results[0].Address
	concert.Coordinates[0] = data.Results[0].Geometry.Location.Latitude
	concert.Coordinates[1] = data.Results[0].Geometry.Location.Longtitude
	return true
}
