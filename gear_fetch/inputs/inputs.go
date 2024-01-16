package inputs

import (
	"bufio"
	"embed"
	"log"
	"strings"
)

//go:embed data/farming_locations.csv
var gearFile embed.FS
var gearInfoList []GearInfo
var gearInfoByName map[string]GearInfo

type GearInfo struct {
	Name     string
	Location string
}

func init() {
	file, err := gearFile.Open("data/farming_locations.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gearInfoList = make([]GearInfo, 0)
	gearInfoByName = make(map[string]GearInfo)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		components := strings.Split(scanner.Text(), ",")
		newGearInfo := GearInfo{
			Name:     components[0],
			Location: components[1],
		}
		gearInfoList = append(gearInfoList, newGearInfo)
		gearInfoByName[newGearInfo.Name] = newGearInfo
	}
}

func SortedGearInfo() []GearInfo {
	// maintain order from file
	return gearInfoList
}

func GearInfoByName() map[string]GearInfo {
	return gearInfoByName
}
