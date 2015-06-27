package utils

import (
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type FileGenerator struct {
	Location string
}

// CreateFile creates a file with the current nano seconds as the filename
// and returns this time stamp (i.e. filename)
func (f FileGenerator) CreateFile() string {
	filename := strconv.Itoa(time.Now().Nanosecond())
	file, _ := os.Create(filepath.Join(f.Location, filename))
	file.Close()
	return filename
}
