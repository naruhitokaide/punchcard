package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateFile(t *testing.T) {
	testDir := "testDir"
	os.MkdirAll(testDir, 0755)
	filegen := RandomFileGenerator{testDir}
	filename := filegen.CreateFile()
	if _, err := os.Stat(filepath.Join(testDir, filename)); os.IsNotExist(err) {
		t.Errorf("Expected file (%s) to be created", filename)
	}
	os.RemoveAll(testDir)
}
