package api

import "time"

type Artists []Artist

type Artist struct {
	Id                  int
	Name                string
	Image               string
	Members             []string
	CreationDate        int
	FirstAlbum          time.Time
	FirstAlbumFormatted string // timeLayout variant of FirstAlbum
	Concerts            []Concert
}

type ShortArtists []ShortArtist

type ShortArtist struct {
	Id    int
	Name  string
	Image string
}

type Concert struct {
	Location      string
	Coordinates   [2]float64 // latitude first element, longtitude second element
	Date          []time.Time
	DateFormatted []string // timeLayout variant of Date
}

type Suggestions []Suggestion

type Suggestion struct {
	Name string
	Type string
}

type FilterParams struct {
	Search string `json:"Search"`
	Filter struct {
		IsEnabled    bool `json:"IsEnabled"`
		CreationDate struct {
			IsEnabled  bool `json:"IsEnabled"`
			FromYear   int  `json:"FromYear"`
			BeforeYear int  `json:"BeforeYear"`
		} `json:"CreationDate"`
		Cities struct {
			IsEnabled bool     `json:"IsEnabled"`
			Locations []string `json:"Locations"`
		} `json:"Cities"`
		FirstAlbumDate struct {
			IsEnabled  bool   `json:"IsEnabled"`
			FromDate   string `json:"FromDate"`
			BeforeDate string `json:"BeforeDate"`
		} `json:"FirstAlbumDate"`
		CountMembers struct {
			IsEnabled bool `json:"IsEnabled"`
			From      int  `json:"From"`
			Before    int  `json:"Before"`
		} `json:"CountMembers"`
	} `json:"Filter"`
}

// Artists and Relation - API JSON formats
type artistsJSON struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type relationsJSON struct {
	Index []struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

// Geocoding - Google MAP API JSON format
type geocodingJSON struct {
	Results []struct {
		Address  string `json:"formatted_address"`
		Geometry struct {
			Location struct {
				Latitude   float64 `json:"lat"`
				Longtitude float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}
