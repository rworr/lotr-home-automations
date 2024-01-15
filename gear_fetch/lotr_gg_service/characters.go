package lotr_gg_service

// Fetch all supported characters shown on lotr.gg

import (
	"fmt"

	"golang.org/x/net/html"
)

const (
	characterPageUrl = "https://lotr.gg/characters/"
	gridCellTag      = "unit-card-grid__cell"
	unitCardTag      = "unit-card__name"
)

func GetCharacters() (map[string]string, error) {
	body, err := callWebsite(characterPageUrl)
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
		if attr.Key == "class" && attr.Val == gridCellTag {
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
		if attr.Key == "class" && attr.Val == unitCardTag {
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
