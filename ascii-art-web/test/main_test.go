package test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

const HostURL = "http://localhost:8080/"

func TestIndexPage(t *testing.T) {
	// Test For Get
	res, err := http.Get(HostURL)
	if err != nil {
		t.Fatal(err)
	}
	expectedStatus := http.StatusOK
	if res.StatusCode != expectedStatus {
		t.Errorf("For %q http.Get returns %d instead of %d", HostURL, res.StatusCode, expectedStatus)
	}

	// Test For Post
	res, err = http.Post(HostURL, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	expectedStatus = http.StatusMethodNotAllowed
	if res.StatusCode != expectedStatus {
		t.Errorf("For %q http.Post returns %d instead of %d", HostURL, res.StatusCode, expectedStatus)
	}

	// Test For 404
	res, err = http.Get(HostURL + "/dasda")
	if err != nil {
		t.Fatal(err)
	}
	expectedStatus = http.StatusNotFound
	if res.StatusCode != expectedStatus {
		t.Errorf("For %q http.Get returns %d instead of %d", HostURL, res.StatusCode, expectedStatus)
	}
}

func TestAsciiArtPage(t *testing.T) {
	hostUrl := fmt.Sprintf("%v%v", HostURL, "ascii-art")
	// Test For Post
	post := url.Values{
		"input":  []string{"Hello"},
		"text":   []string{"Hello"},
		"font":   []string{"standard"},
		"theme":  []string{"standard"},
		"submit": []string{"show"},
		"type":   []string{"show"},
	}
	res, err := http.PostForm(hostUrl, post)
	if err != nil {
		t.Fatal(err)
	}
	expectedStatus := http.StatusOK
	if res.StatusCode != expectedStatus {
		t.Errorf("For %q http.Post returns %d instead of %d", hostUrl, res.StatusCode, expectedStatus)
	}

	// For Bad Request
	res, err = http.PostForm(hostUrl, nil)
	if err != nil {
		t.Fatal(err)
	}
	expectedStatus = http.StatusBadRequest
	if res.StatusCode != expectedStatus {
		t.Errorf("For %q http.Post returns %d instead of %d", hostUrl, res.StatusCode, expectedStatus)
	}

	// For Bad Request
	post = url.Values{
		"input":  []string{"ффф"},
		"text":   []string{"ффф"},
		"font":   []string{"standard"},
		"theme":  []string{"standard"},
		"submit": []string{"show"},
		"type":   []string{"show"},
	}
	res, err = http.PostForm(hostUrl, post)
	if err != nil {
		t.Fatal(err)
	}
	expectedStatus = http.StatusBadRequest
	if res.StatusCode != expectedStatus {
		t.Errorf("For %q http.Post returns %d instead of %d", hostUrl, res.StatusCode, expectedStatus)
	}

	// Test For Get
	res, err = http.Get(hostUrl)
	if err != nil {
		t.Fatal(err)
	}
	expectedStatus = http.StatusMethodNotAllowed
	if res.StatusCode != expectedStatus {
		t.Errorf("For %q http.Get returns %d instead of %d", HostURL, res.StatusCode, expectedStatus)
	}

	// Test For 404
	res, err = http.Get(hostUrl + "/dasda")
	if err != nil {
		t.Fatal(err)
	}
	expectedStatus = http.StatusNotFound
	if res.StatusCode != expectedStatus {
		t.Errorf("For %q http.Get returns %d instead of %d", HostURL, res.StatusCode, expectedStatus)
	}
}
