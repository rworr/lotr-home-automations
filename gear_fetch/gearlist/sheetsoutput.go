package gearlist

import (
	"fmt"
	"gear-fetch/inputs"
	"log"
	"strconv"
	"strings"

	"google.golang.org/api/sheets/v4"
)

const sheetName = "Gear Output"

func OutputGearListToSheets(gearList GearList, homeSheet *HomeSpreadsheet) {
	service := homeSheet.Service.Sheets
	addNewGearSheet(homeSheet)
	sortedCharacterLevels := getSortedCharacterLevels(gearList)
	sortedGear := inputs.SortedGearInfo()

	numCols := 2 + len(sortedCharacterLevels)
	numRows := 2 + len(sortedGear)

	values := make([][]interface{}, numRows)
	for i := range values {
		values[i] = make([]interface{}, numCols)
	}

	getSheetsOutputHeaders(values, sortedCharacterLevels)
	for i, gearInfo := range sortedGear {
		characterGear := gearList[gearInfo]
		getSheetsOutputString(values[i+2], gearInfo, characterGear, sortedCharacterLevels)
	}

	sheetRange := getRange(numCols, numRows)
	valueRange := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Range:          sheetRange,
		Values:         values,
	}
	_, err := service.Spreadsheets.Values.Update(homeSheet.Sheet.SpreadsheetId, sheetRange, valueRange).
		ValueInputOption("USER_ENTERED").IncludeValuesInResponse(false).Do()
	if err != nil {
		log.Fatalf("Unable to write to sheet: %v", err)
	}
}

func addNewGearSheet(homeSheet *HomeSpreadsheet) {
	service := homeSheet.Service.Sheets
	props := &sheets.SheetProperties{Title: sheetName}
	req := sheets.Request{AddSheet: &sheets.AddSheetRequest{Properties: props}}
	batchRequest := &sheets.BatchUpdateSpreadsheetRequest{Requests: []*sheets.Request{&req}}
	_, err := service.Spreadsheets.BatchUpdate(homeSheet.Sheet.SpreadsheetId, batchRequest).Do()
	if err != nil {
		log.Fatalf("Unable to add new sheet: %v", err)
	}
}

func getSheetsOutputHeaders(values [][]interface{}, characterGearLevels []CharacterGearLevel) {
	for i := 0; i < len(characterGearLevels); i++ {
		values[0][i+2] = characterGearLevels[i].Character
		values[1][i+2] = characterGearLevels[i].Level
	}
}

func getSheetsOutputString(
	values []interface{},
	info inputs.GearInfo,
	entry characterGearEntry,
	sortedCharacters []CharacterGearLevel,
) {
	values[0] = info.Name
	values[1] = info.Location
	for i, char := range sortedCharacters {
		values[i+2] = strconv.Itoa(entry[char])
	}
}

func getRange(numCols, numRows int) string {
	return fmt.Sprintf("%v!A1:%v%v", sheetName, base26(numCols), numRows)
}

func base26(n int) string {
	builder := strings.Builder{}
	for i := n; i > 0; {
		x := (i - 1) % 26
		builder.WriteByte(byte(x + 65))
		i = (i - x) / 26
	}

	r := []rune(builder.String())
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}
