package utils

import (
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// FileGenerator is an interface that provides a method to generate files.
type FileGenerator interface {
	CreateFile() (string, error)
}

// RandomFileGenerator holds the information of the Location of a file.
type RandomFileGenerator struct {
	Location string
}

// CreateFile creates a file with the current nano seconds as the filename in
// the file generators location and returns the time stamp (i.e. filename).
func (f RandomFileGenerator) CreateFile() (filename string, err error) {
	filename = strconv.Itoa(time.Now().Nanosecond())
	file, err := os.Create(filepath.Join(f.Location, filename))
	if err != nil {
		return
	}
	err = file.Close()
	return
}
