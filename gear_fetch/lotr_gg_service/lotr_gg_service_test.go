package lotr_gg_service_test

import (
	gearfetch "gear-fetch/lotr_gg_service"
	"testing"
)

func TestGetCharacters(t *testing.T) {
	characterMap, err := gearfetch.GetCharacters()
	if err != nil {
		t.Error(err)
	}

	var aeldred = characterMap["Aeldred"]
	var expectedUrl = "/characters/unit_rohan_healer_01_pvp/"

	if aeldred != expectedUrl {
		t.Errorf("Expected: %v, found %v", expectedUrl, aeldred)
	}
}

func TestGetGear(t *testing.T) {
	characterMap := make(gearfetch.CharacterUrls)
	characterMap["Aeldred"] = "/characters/unit_rohan_healer_01_pvp/"

	gearLevels, err := gearfetch.GetCharacterGear("Aeldred", characterMap)
	if err != nil {
		t.Error(err)
	}

	for _, level := range gearLevels {
		if level.Level == "Gear Level 7" && level.Gear["Crest II"] != 6 {
			t.Errorf("Expected Gear Level 7 to require 6 Crest II, found gear: %v", level.Gear)
		}
	}
}
