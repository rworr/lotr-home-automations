package main

import (
	"fmt"
	"gear-fetch/lotr_gg_service"
)

func main() {
	characters, err := lotr_gg_service.GetCharacters()
	if err != nil {
		return
	}

	for char, url := range characters {
		fmt.Printf("%v - %v\n", char, url)
	}

	gearLevels, err := lotr_gg_service.GetCharacterGear("Aeldred", characters)
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
