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
