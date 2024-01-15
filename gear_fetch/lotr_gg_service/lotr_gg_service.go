package lotr_gg_service

/*
	Wrapper around scraping information from lotr.gg
*/

import (
	"fmt"
	"io"

	http "github.com/saucesteals/fhttp"
	"github.com/saucesteals/mimic"
)

type WebsiteError struct {
	statusCode int
}

func (err WebsiteError) Error() string {
	return fmt.Sprintf("Received error code %3d", err.statusCode)
}

func callWebsite(url string) (io.Reader, error) {
	latestVersion := mimic.MustGetLatestVersion(mimic.PlatformWindows)
	clientSpec, _ := mimic.Chromium(mimic.BrandChrome, latestVersion)

	client := &http.Client{
		Transport: clientSpec.ConfigureTransport(&http.Transport{}),
	}

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:61.0) Gecko/20100101 Firefox/61.0")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Accept-Language", "en-US,en;q=0.5")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, WebsiteError{statusCode: response.StatusCode}
	}

	return response.Body, nil
}
