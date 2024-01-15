package gear_fectch

/*
	Wrapper around scraping information from lotr.gg
*/

import (
	"fmt"
	"io"

	"golang.org/x/net/html"

	http "github.com/saucesteals/fhttp"
	"github.com/saucesteals/mimic"
)

type WebsiteError struct {
	statusCode int
}

func (err WebsiteError) Error() string {
	return fmt.Sprintf("Received error code %3d", err.statusCode)
}

func GetCharacters() (map[string]string, error) {
	body, err := callWebsite("https://lotr.gg/characters/")
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}

	characterMap := make(map[string]string)
	crawlForCharacters(doc, characterMap)
	return characterMap, nil
}

func crawlForCharacters(node *html.Node, characterUrls map[string]string) {
	var enteringGridCell bool
	for _, attr := range node.Attr {
		if attr.Key == "class" && attr.Val == "unit-card-grid__cell" {
			enteringGridCell = true
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if enteringGridCell {
			crawlForCharacterData(child, characterUrls)
		} else {
			crawlForCharacters(child, characterUrls)
		}
	}
}

func crawlForCharacterData(node *html.Node, characterUrls map[string]string) {
	if node.Data != "a" {
		return
	}

	var characterUrl string
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			characterUrl = attr.Val
		}
	}

	var characterName string
	for child := node.FirstChild; child != nil && characterName == ""; child = child.NextSibling {
		characterName = crawlForCharacterCard(child, characterUrls)
	}

	if characterName == "" || characterUrl == "" {
		panic(fmt.Sprintf("Malformed HTML - unable to get name (%v) + link (%v)", characterName, characterUrl))
	}

	characterUrls[characterName] = characterUrl
}

func crawlForCharacterCard(node *html.Node, characterUrls map[string]string) string {
	var enteringUnitCard bool
	for _, attr := range node.Attr {
		if attr.Key == "class" && attr.Val == "unit-card__name" {
			enteringUnitCard = true
		}
	}

	var characterName string
	for child := node.FirstChild; child != nil && characterName == ""; child = child.NextSibling {
		if enteringUnitCard && child.Type == html.TextNode {
			characterName = child.Data
		} else {
			characterName = crawlForCharacterCard(child, characterUrls)
		}
	}

	return characterName
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
