package gearlist_test

import (
	"bufio"
	"gear-fetch/gearlist"
	"gear-fetch/inputs"
	"gear-fetch/lotr_gg_service"
	"strings"
	"testing"
)

func TestNewGearList(t *testing.T) {
	characterGear := map[string]lotr_gg_service.GearLevels{"Golburz": {lotr_gg_service.GearLevel{Level: "Gear Level 1", Gear: lotr_gg_service.GearMap{"Glass Vial I": 1, "Hide I": 1, "Lamp I": 1, "Parchment I": 1, "Satchel I": 1, "Worn Ring I": 1}}, lotr_gg_service.GearLevel{Level: "Gear Level 2", Gear: lotr_gg_service.GearMap{"Glass Vial I": 1, "Hide I": 1, "Lamp I": 1, "Parchment I": 1, "Resin I": 1, "Satchel I": 1, "Thread I": 3, "Worn Ring I": 1}}, lotr_gg_service.GearLevel{Level: "Gear Level 3", Gear: lotr_gg_service.GearMap{"Blessing I": 2, "Glass Vial I": 1, "Hide I": 1, "Lamp I": 1, "Parchment I": 2, "Resin I": 1, "Satchel I": 1, "Thread I": 5, "Worn Ring I": 1}}, lotr_gg_service.GearLevel{Level: "Gear Level 4", Gear: lotr_gg_service.GearMap{"Armor Essence II": 4, "Blessing I": 2, "Crest II": 2, "Damage Essence II": 2, "Hide I": 1, "Resist Essence II": 4, "Shadow Crystal II": 20, "Tabard II": 6, "Thread I": 2, "Worn Ring I": 1}}, lotr_gg_service.GearLevel{Level: "Gear Level 5", Gear: lotr_gg_service.GearMap{"Armor Essence II": 6, "Cotton II": 9, "Crest II": 2, "Damage Essence II": 4, "Focus Essence II": 2, "Health Essence II": 2, "Resist Essence II": 6, "Shadow Crystal II": 40, "Tabard II": 10}}, lotr_gg_service.GearLevel{Level: "Gear Level 6", Gear: lotr_gg_service.GearMap{"Armor Essence II": 10, "Blessing II": 6, "Cotton II": 9, "Crest II": 4, "Damage Essence II": 6, "Focus Essence II": 4, "Health Essence II": 4, "Jasper II": 12, "Resist Essence II": 8, "Shadow Crystal II": 64, "Tabard II": 10}}, lotr_gg_service.GearLevel{Level: "Gear Level 7", Gear: lotr_gg_service.GearMap{"Armor Essence II": 12, "Blessing II": 6, "Cotton II": 15, "Crest II": 6, "Damage Essence II": 8, "Focus Essence II": 6, "Health Essence II": 6, "Jasper II": 30, "Resist Essence II": 8, "Shadow Crystal II": 80, "Tabard II": 14}}, lotr_gg_service.GearLevel{Level: "Gear Level 8", Gear: lotr_gg_service.GearMap{"Armor Essence III": 4, "Blessing II": 6, "Clay III": 2, "Cotton III": 15, "Crest III": 3, "Damage Essence III": 4, "Focus Essence II": 6, "Health Essence II": 6, "Jasper II": 12, "Linen III": 6, "Resist Essence III": 8, "Shadow Crystal Fragments III": 80, "Shadow Crystal II": 24, "Tabard II": 4, "Tabard III": 9, "Tongs II": 8}}, lotr_gg_service.GearLevel{Level: "Gear Level 9", Gear: lotr_gg_service.GearMap{"Armor Essence III": 12, "Clay III": 2, "Cotton III": 10, "Crest III": 3, "Damage Essence III": 12, "Focus Essence III": 4, "Health Essence III": 4, "Jasper III": 9, "Linen III": 25, "Resist Essence III": 20, "Shadow Crystal Fragments III": 260, "Tabard III": 24, "Tongs II": 10}}, lotr_gg_service.GearLevel{Level: "Gear Level 10", Gear: lotr_gg_service.GearMap{}}}}
	golburzGear := gearlist.ParseGearLevels(characterGear)

	gearInfo := inputs.GearInfo{Name: "Tabard III", Location: " 6-1"}
	characterLevel := gearlist.CharacterGearLevel{Character: "Golburz", Level: "Gear Level 9"}
	if quantity := golburzGear[gearInfo][characterLevel]; quantity != 24 {
		t.Errorf("Expected quantity 4, found %v\n", quantity)
	}
}

func TestOutputGearList(t *testing.T) {
	characterGear := map[string]lotr_gg_service.GearLevels{"Golburz": {lotr_gg_service.GearLevel{Level: "Gear Level 1", Gear: lotr_gg_service.GearMap{"Glass Vial I": 1, "Hide I": 1, "Lamp I": 1, "Parchment I": 1, "Satchel I": 1, "Worn Ring I": 1}}, lotr_gg_service.GearLevel{Level: "Gear Level 2", Gear: lotr_gg_service.GearMap{"Glass Vial I": 1, "Hide I": 1, "Lamp I": 1, "Parchment I": 1, "Resin I": 1, "Satchel I": 1, "Thread I": 3, "Worn Ring I": 1}}, lotr_gg_service.GearLevel{Level: "Gear Level 3", Gear: lotr_gg_service.GearMap{"Blessing I": 2, "Glass Vial I": 1, "Hide I": 1, "Lamp I": 1, "Parchment I": 2, "Resin I": 1, "Satchel I": 1, "Thread I": 5, "Worn Ring I": 1}}, lotr_gg_service.GearLevel{Level: "Gear Level 4", Gear: lotr_gg_service.GearMap{"Armor Essence II": 4, "Blessing I": 2, "Crest II": 2, "Damage Essence II": 2, "Hide I": 1, "Resist Essence II": 4, "Shadow Crystal II": 20, "Tabard II": 6, "Thread I": 2, "Worn Ring I": 1}}, lotr_gg_service.GearLevel{Level: "Gear Level 5", Gear: lotr_gg_service.GearMap{"Armor Essence II": 6, "Cotton II": 9, "Crest II": 2, "Damage Essence II": 4, "Focus Essence II": 2, "Health Essence II": 2, "Resist Essence II": 6, "Shadow Crystal II": 40, "Tabard II": 10}}, lotr_gg_service.GearLevel{Level: "Gear Level 6", Gear: lotr_gg_service.GearMap{"Armor Essence II": 10, "Blessing II": 6, "Cotton II": 9, "Crest II": 4, "Damage Essence II": 6, "Focus Essence II": 4, "Health Essence II": 4, "Jasper II": 12, "Resist Essence II": 8, "Shadow Crystal II": 64, "Tabard II": 10}}, lotr_gg_service.GearLevel{Level: "Gear Level 7", Gear: lotr_gg_service.GearMap{"Armor Essence II": 12, "Blessing II": 6, "Cotton II": 15, "Crest II": 6, "Damage Essence II": 8, "Focus Essence II": 6, "Health Essence II": 6, "Jasper II": 30, "Resist Essence II": 8, "Shadow Crystal II": 80, "Tabard II": 14}}, lotr_gg_service.GearLevel{Level: "Gear Level 8", Gear: lotr_gg_service.GearMap{"Armor Essence III": 4, "Blessing II": 6, "Clay III": 2, "Cotton III": 15, "Crest III": 3, "Damage Essence III": 4, "Focus Essence II": 6, "Health Essence II": 6, "Jasper II": 12, "Linen III": 6, "Resist Essence III": 8, "Shadow Crystal Fragments III": 80, "Shadow Crystal II": 24, "Tabard II": 4, "Tabard III": 9, "Tongs II": 8}}, lotr_gg_service.GearLevel{Level: "Gear Level 9", Gear: lotr_gg_service.GearMap{"Armor Essence III": 12, "Clay III": 2, "Cotton III": 10, "Crest III": 3, "Damage Essence III": 12, "Focus Essence III": 4, "Health Essence III": 4, "Jasper III": 9, "Linen III": 25, "Resist Essence III": 20, "Shadow Crystal Fragments III": 260, "Tabard III": 24, "Tongs II": 10}}, lotr_gg_service.GearLevel{Level: "Gear Level 10", Gear: lotr_gg_service.GearMap{}}}}
	golburzGear := gearlist.ParseGearLevels(characterGear)

	builder := strings.Builder{}
	bufioWriter := bufio.NewWriter(&builder)
	gearlist.OutputGearListToCSV(bufioWriter, golburzGear)
	bufioWriter.Flush()

	gearOutput := builder.String()
	expectedSubstrings := [3]string{
		",,Golburz",
		"Gear Level 3,Gear Level 4,Gear Level 5",
		"Resin I, 1-7,0,1,1,0,0,0,0,0,0",
	}

	for _, substring := range expectedSubstrings {
		if !strings.Contains(gearOutput, substring) {
			t.Errorf("Missing %v in output\n\n%v", substring, gearOutput)
		}
	}
}

func TestGetHomeSpreadsheet(t *testing.T) {
	sheet := gearlist.GetHomeSpreadsheet()
	readRange := "Farming Plan!A2:E2"

	resp, err := sheet.Sheets.Spreadsheets.Values.Get(sheet.Sheet.SpreadsheetId, readRange).Do()
	if err != nil {
		t.Errorf("Unable to retrieve data from sheet: %v", err)
	}

	if resp.Values[0][0] != "Farming" {
		t.Errorf("Expected 'Farming' as return value, got :%#v", resp.Values)
	}
}
