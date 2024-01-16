package gearlist

import (
	_ "bufio"
	_ "fmt"
	"os"
	"path/filepath"
	"runtime"
)

func OutputToFile(gearList GearList) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("panic!!")
	}

	outputFileName := filepath.Join(filepath.Dir(filename), "gearlist.csv")
	outfile, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	outfile.WriteString("Test")
}
