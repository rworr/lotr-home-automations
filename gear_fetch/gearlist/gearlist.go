package gearlist

import (
	"bufio"
	"embed"
	"log"
	"strings"

	"gear-fetch/lotr_gg_service"
)

//go:embed data/farming_locations.csv
var gearFile embed.FS
var gearInfoList []GearInfo
var gearInfoByName map[string]GearInfo

type GearInfo struct {
	Name     string
	Location string
}

type CharacterGearLevel struct {
	Character string
	Level     string
}

type characterGearEntry map[CharacterGearLevel]int
type GearList map[GearInfo]characterGearEntry

func init() {
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

func SortedGearInfo() []GearInfo {
	// maintain order from file
	return gearInfoList
}

func NewGearList(characterGear map[string]lotr_gg_service.GearLevels) GearList {
	gearList := make(GearList, len(gearInfoList))
	for _, gear := range gearInfoList {
		gearList[gear] = make(characterGearEntry)
	}

	for character, gearLevels := range characterGear {
		for _, gearLevel := range gearLevels {
			characterLevel := CharacterGearLevel{
				Character: character,
				Level:     gearLevel.Level,
			}

			for gearName, quantity := range gearLevel.Gear {
				gear := gearInfoByName[gearName]
				if gearList[gear] != nil {
					gearList[gear][characterLevel] = quantity
				} else {
					unknownGearInfo := GearInfo{
						Name:     gearName,
						Location: "unknown",
					}
					if gearList[unknownGearInfo] == nil {
						gearList[unknownGearInfo] = make(characterGearEntry)
					}

					gearList[unknownGearInfo][characterLevel] = quantity
				}

			}
		}
	}

	return gearList
}
