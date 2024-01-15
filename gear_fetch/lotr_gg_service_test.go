package gear_fectch_test

import (
	gearfetch "gear-fetch"
	"testing"
)

/*
func TestCallWebsite(t *testing.T) {
	status, err := gearfetch.CallWebsite(aeldred)
	if err != nil {
		t.Error(err)
	}

	t.Log(status)
}
*/

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
