package controllers

import (
	"github.com/felipe-michelon/url-shortener/database"
	"github.com/felipe-michelon/url-shortener/models"

	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExistingUrl(t *testing.T) {
	url := models.Url{Shortened: "shortened", Original: "http://original.com"}

	database.DB.Create(&url)
	defer database.DB.Unscoped().Delete(&url)

	ts := httptest.NewServer(SetupRouter())
	defer ts.Close()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}

	req, _ := client.Get(fmt.Sprintf("%s/%s", ts.URL, url.Shortened))

	assert.Equal(t, 301, req.StatusCode)
	assert.Equal(t, req.Header.Get("Location"), url.Original)
}

func TestGetNonExistingUrl(t *testing.T) {
	ts := httptest.NewServer(SetupRouter())
	defer ts.Close()

	req, _ := http.Get(fmt.Sprintf("%s/shortened", ts.URL))
	body, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	assert.Equal(t, 404, req.StatusCode)
	assert.Equal(t, "Url not found", string(body))
}
