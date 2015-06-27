package utils

import (
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type FileGenerator interface {
	CreateFile() string
}

type RandomFileGenerator struct {
	Location string
}

// CreateFile creates a file with the current nano seconds as the filename in
// the file generators location and returns the time stamp (i.e. filename).
func (f RandomFileGenerator) CreateFile() string {
	filename := strconv.Itoa(time.Now().Nanosecond())
	file, _ := os.Create(filepath.Join(f.Location, filename))
	file.Close()
	return filename
}
