package api

var (
	AllArtists     *Artists
	NewAllArtists  *Artists
	AllSuggestions *Suggestions
	AllCities      *[]string
)

// API urls to use
const (
	artistsAPI  = "https://groupietrackers.herokuapp.com/api/artists"
	relationAPI = "https://groupietrackers.herokuapp.com/api/relation"
)

// Layout for time parsing
const (
	timeLayout          = "02-01-2006"
	timeLayoutForClient = "2006-01-02T00:00:00.000Z"
)

// Google MAP API params
const (
	googleMapAPI = "AIzaSyClp0He8MpMVMaSxh7WU0gRHmTi4j2PB5U"
	geocodingURL = "https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s" // by the city name
)

// Cache path
const cachePath = "api/cache/all.json"
