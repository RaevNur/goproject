package api

import (
	"strconv"
	"strings"
	"time"
)

func CheckFilterFields(params FilterParams) bool {
	if !params.Filter.IsEnabled {
		return true
	}
	if params.Filter.FirstAlbumDate.IsEnabled {
		from, err := time.Parse(timeLayoutForClient, params.Filter.FirstAlbumDate.FromDate)
		if err != nil {
			return false
		}
		before, err := time.Parse(timeLayoutForClient, params.Filter.FirstAlbumDate.BeforeDate)
		if err != nil {
			return false
		}
		if from.After(before) {
			return false
		}
	}
	if params.Filter.CreationDate.IsEnabled && (params.Filter.CreationDate.FromYear < 0 || params.Filter.CreationDate.BeforeYear < 0) || (params.Filter.CreationDate.FromYear > params.Filter.CreationDate.BeforeYear) {
		return false
	}
	if params.Filter.CountMembers.IsEnabled && (params.Filter.CountMembers.From < 0 || params.Filter.CountMembers.Before < 0) || (params.Filter.CountMembers.From > params.Filter.CountMembers.Before) {
		return false
	}
	return true
}

func FilterByParams(params FilterParams) *Artists {
	filtered := &Artists{}
	searchOn, filterOn := false, params.Filter.IsEnabled
	if len(params.Search) > 0 {
		searchOn = true
	}
	if !searchOn && !filterOn {
		return AllArtists
	}
	for _, artist := range *AllArtists {
		if searchOn && searchResult(artist, params.Search) {
			if !filterOn || (filterOn && filterResult(artist, params)) {
				*filtered = append(*filtered, artist)
			}
		} else if !searchOn && filterResult(artist, params) {
			*filtered = append(*filtered, artist)
		}
	}
	return filtered
}

func searchResult(artist Artist, looking string) bool {
	if strings.Contains(artist.Name, looking) {
		return true
	}
	for _, member := range artist.Members {
		if strings.Contains(member, looking) {
			return true
		}
	}
	if strings.Contains(strconv.Itoa(artist.CreationDate), looking) {
		return true
	}
	if strings.Contains(artist.FirstAlbumFormatted, looking) {
		return true
	}
	for _, concert := range artist.Concerts {
		if strings.Contains(concert.Location, looking) {
			return true
		}
		for _, date := range concert.DateFormatted {
			if strings.Contains(date, looking) {
				return true
			}
		}
	}
	return false
}

func filterResult(artist Artist, filter FilterParams) bool {
	if filter.Filter.CreationDate.IsEnabled {
		if !(filter.Filter.CreationDate.FromYear <= artist.CreationDate &&
			artist.CreationDate <= filter.Filter.CreationDate.BeforeYear) {
			return false
		}
	}
	if filter.Filter.FirstAlbumDate.IsEnabled {
		// from, before := filter.Filter.FirstAlbumDate.FromDate, filter.Filter.FirstAlbumDate.BeforeDate
		from, _ := time.Parse(timeLayoutForClient, filter.Filter.FirstAlbumDate.FromDate)
		before, _ := time.Parse(timeLayoutForClient, filter.Filter.FirstAlbumDate.BeforeDate)
		if !(artist.FirstAlbum.After(from) && artist.FirstAlbum.Before(before)) {
			return false
		}
	}
	if filter.Filter.CountMembers.IsEnabled {
		members := len(artist.Members)
		if !(filter.Filter.CountMembers.From <= members &&
			members <= filter.Filter.CountMembers.Before) {
			return false
		}
	}
	if filter.Filter.Cities.IsEnabled {
		success := false
		for _, city := range filter.Filter.Cities.Locations {
			for _, concert := range artist.Concerts {
				if city == concert.Location {
					success = true
					break
				}
			}
			if success {
				break
			}
		}
		if !success {
			return false
		}
	}
	return true
}
