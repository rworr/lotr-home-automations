package gearlist

import (
	"bufio"
	"gear-fetch/inputs"
	"sort"
	"strconv"
	"strings"
)

const sep = ","

func OutputGearListToCSV(writer *bufio.Writer, gearList GearList) {
	sortedCharacterLevels := getSortedCharacterLevels(gearList)
	sortedGear := inputs.SortedGearInfo()

	writer.WriteString(getCSVOutputHeaders(sortedCharacterLevels))

	for _, gearInfo := range sortedGear {
		characterGear := gearList[gearInfo]
		writer.WriteString(getCSVOutputString(gearInfo, characterGear, sortedCharacterLevels))
	}
}

func getSortedCharacterLevels(gearList GearList) []CharacterGearLevel {
	characterGearLevelSet := make(map[CharacterGearLevel]bool)

	for _, gearEntry := range gearList {
		for characterGearLevel, _ := range gearEntry {
			characterGearLevelSet[characterGearLevel] = true
		}
	}

	characterGearLevels := make([]CharacterGearLevel, 0, len(characterGearLevelSet))
	for characterGear := range characterGearLevelSet {
		characterGearLevels = append(characterGearLevels, characterGear)
	}

	// sort on character name first then level name second
	sort.Slice(characterGearLevels, func(i, j int) bool {
		if characterGearLevels[i].Character == characterGearLevels[j].Character {
			return characterGearLevels[i].Level < characterGearLevels[j].Level
		}
		return characterGearLevels[i].Character < characterGearLevels[j].Character
	})

	return characterGearLevels
}

func getCSVOutputHeaders(characterGearLevels []CharacterGearLevel) string {
	builder := strings.Builder{}
	builder.WriteString(sep)
	builder.WriteString(sep)
	builder.WriteString(characterGearLevels[0].Character)
	for _, s := range characterGearLevels[1:] {
		builder.WriteString(sep)
		builder.WriteString(s.Character)
	}
	builder.WriteString("\n")

	builder.WriteString(sep)
	builder.WriteString(sep)
	builder.WriteString(characterGearLevels[0].Level)
	for _, s := range characterGearLevels[1:] {
		builder.WriteString(sep)
		builder.WriteString(s.Level)
	}
	builder.WriteString("\n")

	return builder.String()
}

func getCSVOutputString(info inputs.GearInfo, entry characterGearEntry, sortedCharacters []CharacterGearLevel) string {
	builder := strings.Builder{}
	builder.WriteString(info.Name)
	builder.WriteString(sep)
	builder.WriteString(info.Location)
	builder.WriteString(sep)

	builder.WriteString(strconv.Itoa(entry[sortedCharacters[0]]))
	for _, character := range sortedCharacters[1:] {
		builder.WriteString(sep)
		builder.WriteString(strconv.Itoa(entry[character]))
	}
	builder.WriteString("\n")

	return builder.String()
}
