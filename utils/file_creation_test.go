package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateFile(t *testing.T) {
	testDir := "testDir"
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Errorf("Failed to create test directory: %v", err)
	}
	filegen := RandomFileGenerator{testDir}
	filename, err := filegen.CreateFile()
	if err != nil {
		t.Errorf("Failed to create file: %v", err)
	}
	if _, err := os.Stat(filepath.Join(testDir, filename)); os.IsNotExist(err) {
		t.Errorf("Expected file (%s) to be created", filename)
	}
	err = os.RemoveAll(testDir)
	if err != nil {
		t.Errorf("Failed to delete test directory: %v", err)
	}
}
