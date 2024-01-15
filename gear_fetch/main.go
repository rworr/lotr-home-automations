package main

import (
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

	for _, level := range gearLevels {
		println(level.Level)
		for name, qty := range level.Gear {
			println(name, qty)
		}
		println()
	}
}
