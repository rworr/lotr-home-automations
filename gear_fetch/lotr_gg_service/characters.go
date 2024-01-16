package lotr_gg_service

// Fetch all supported characters shown on lotr.gg

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

const (
	characterPageUrl = "https://lotr.gg/characters/"
	gridCellTag      = "unit-card-grid__cell"
	unitCardTag      = "unit-card__name"
)

type CharacterUrls map[string]string

type FileNameError struct {
	pc   uintptr
	file string
	line int
}

func (err FileNameError) Error() string {
	return fmt.Sprintf(
		"Unable to determine root file path, got pc: %v, file: %v, line: %v",
		err.pc,
		err.file,
		err.line,
	)
}

func GetCharacters() (CharacterUrls, error) {
	body, err := callWebsite(characterPageUrl)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}

	characterMap := make(CharacterUrls)
	crawlForCharacters(doc, characterMap)
	return characterMap, nil
}

func DumpCharacters(characterMap CharacterUrls) {
	outputFileName, err := getOutputFileName()
	outfile, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	characters := make([]string, 0, len(characterMap))
	for character := range characterMap {
		characters = append(characters, character)
	}

	sort.Strings(characters)

	outfile.WriteString(strings.Join(characters, "\n"))
}

func crawlForCharacters(node *html.Node, characterUrls CharacterUrls) {
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

func crawlForCharacterData(node *html.Node, characterUrls CharacterUrls) {
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
		characterName = crawlForCharacterCard(child)
	}

	if characterName == "" || characterUrl == "" {
		panic(fmt.Sprintf("Malformed HTML - unable to get name (%v) + link (%v)", characterName, characterUrl))
	}

	characterUrls[characterName] = characterUrl
}

func crawlForCharacterCard(node *html.Node) string {
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
			characterName = crawlForCharacterCard(child)
		}
	}

	return characterName
}

func getOutputFileName() (string, error) {
	pc, filename, line, ok := runtime.Caller(0)
	if !ok {
		return "", FileNameError{pc, filename, line}
	}
	return filepath.Join(filepath.Dir(filename), "../outputs/characters.csv"), nil
}
