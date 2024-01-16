package gearlist

import (
	"gear-fetch/inputs"
	"gear-fetch/lotr_gg_service"
)

type CharacterGearLevel struct {
	Character string
	Level     string
}

type characterGearEntry map[CharacterGearLevel]int
type GearList map[inputs.GearInfo]characterGearEntry

func ParseGearLevels(characterGear map[string]lotr_gg_service.GearLevels) GearList {
	gearInfoList := inputs.SortedGearInfo()
	gearInfoByName := inputs.GearInfoByName()
	gearList := make(GearList, len(gearInfoList))
	for _, gear := range gearInfoList {
		gearList[gear] = make(characterGearEntry)
	}

	for character, gearLevels := range characterGear {
		for _, gearLevel := range gearLevels {
			characterLevel := CharacterGearLevel{
				Character: character,
				Level:     gearLevel.Level,
			}

			for gearName, quantity := range gearLevel.Gear {
				gear := gearInfoByName[gearName]
				if gearList[gear] != nil {
					gearList[gear][characterLevel] = quantity
				} else {
					unknownGearInfo := inputs.GearInfo{
						Name:     gearName,
						Location: "unknown",
					}
					if gearList[unknownGearInfo] == nil {
						gearList[unknownGearInfo] = make(characterGearEntry)
					}

					gearList[unknownGearInfo][characterLevel] = quantity
				}

			}
		}
	}

	return gearList
}
