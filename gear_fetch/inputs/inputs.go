package inputs

import (
	"bufio"
	"embed"
	"log"
	"strings"
)

//go:embed data/farming_locations.csv
var gearFile embed.FS

//go:embed data/input_characters.csv
var characterFile embed.FS

var gearInfoList []GearInfo
var gearInfoByName map[string]GearInfo
var characters []string

type GearInfo struct {
	Name     string
	Location string
}

func init() {
	loadGearInfo()
	loadCharacters()
}

func loadGearInfo() {
	file, err := gearFile.Open("data/farming_locations.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gearInfoList = make([]GearInfo, 0)
	gearInfoByName = make(map[string]GearInfo)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		components := strings.Split(scanner.Text(), ",")
		newGearInfo := GearInfo{
			Name:     components[0],
			Location: components[1],
		}
		gearInfoList = append(gearInfoList, newGearInfo)
		gearInfoByName[newGearInfo.Name] = newGearInfo
	}
}

func loadCharacters() {
	file, err := characterFile.Open("data/input_characters.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	characters = make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		characters = append(characters, scanner.Text())
	}
}

func SortedGearInfo() []GearInfo {
	// maintain order from file
	return gearInfoList
}

func GearInfoByName() map[string]GearInfo {
	return gearInfoByName
}

func Characters() []string {
	return characters
}
