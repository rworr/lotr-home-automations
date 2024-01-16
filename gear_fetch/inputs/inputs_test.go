package inputs

import "testing"

func TestSortedGearInfo(t *testing.T) {
	gear := SortedGearInfo()
	if name := gear[0].Name; name != "Parchment I" {
		t.Errorf("Expected Parchment I, found %v\n", name)
	}

	if name := gear[len(gear)-3].Name; name != "Health Essence III" {
		t.Errorf("Expected Health Essence III, found %v\n", name)
	}
}

func TestCharacters(t *testing.T) {
	chars := Characters()
	if name := chars[0]; name != "Chef Krazhkà" {
		t.Errorf("Expected Parchment I, found %v\n", name)
	}
}
