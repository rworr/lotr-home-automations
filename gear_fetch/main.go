package main

import (
	"fmt"
	"gear-fetch/gearlist"
	"gear-fetch/inputs"
	"gear-fetch/lotr_gg_service"
)

func main() {
	characterMap, err := lotr_gg_service.GetCharacters()
	if err != nil {
		fmt.Printf("Error getting characters: %v", err.Error())
	}
	lotr_gg_service.DumpCharacters(characterMap)

	targetCharacters := inputs.Characters()
	characterGear := make(map[string]lotr_gg_service.GearLevels, len(targetCharacters))
	for _, char := range targetCharacters {
		gearLevels, err := lotr_gg_service.GetCharacterGear(char, characterMap)
		if err != nil {
			fmt.Printf("Error getting character gear for %v: %v", char, err.Error())
		}
		characterGear[char] = gearLevels
	}

	gear := gearlist.ParseGearLevels(characterGear)
	unknownGear := make([]string, 0)
	for gearinfo := range gear {
		if gearinfo.Location == "unknown" {
			unknownGear = append(unknownGear, gearinfo.Name)
		}
	}

	if len(unknownGear) > 0 {
		fmt.Println("Unknown gear found:")
		for _, item := range unknownGear {
			fmt.Println(item)
		}
	}

	homeSheet := gearlist.GetHomeSpreadsheet()
	gearlist.OutputGearListToSheets(gear, homeSheet)
}
