package main

import (
	"gear-fetch/gearlist"
	"gear-fetch/lotr_gg_service"
)

func main() {
	characters, err := lotr_gg_service.GetCharacters()
	if err != nil {
		return
	}

	gearLevels, err := lotr_gg_service.GetCharacterGear("Golburz", characters)
	if err != nil {
		return
	}

	characterGear := make(map[string]lotr_gg_service.GearLevels)
	characterGear["Golburz"] = gearLevels
	list := gearlist.NewGearList(characterGear)

	gearlist.OutputToFile(list)
}
