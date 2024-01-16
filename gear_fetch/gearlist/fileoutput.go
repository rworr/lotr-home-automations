package gearlist

import (
	"bufio"
	_ "bufio"
	"fmt"
	_ "fmt"
	"os"
	"path/filepath"
	"runtime"
)

type FileNameError struct {
	pc   uintptr
	file string
	line int
}

func (err FileNameError) Error() string {
	return fmt.Sprintf(
		"Unable to determine root file path, got pc: %v, file: %v, line: %v",
		err.pc,
		err.file,
		err.line,
	)
}

func OutputToFile(gearList GearList) {
	outputFileName, err := getOutputFileName()
	outfile, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	writer := bufio.NewWriter(outfile)
	defer writer.Flush()

	OutputGearList(writer, gearList)
}

func getOutputFileName() (string, error) {
	pc, filename, line, ok := runtime.Caller(0)
	if !ok {
		return "", FileNameError{pc, filename, line}
	}
	return filepath.Join(filepath.Dir(filename), "../outputs/gearlist.csv"), nil
}
