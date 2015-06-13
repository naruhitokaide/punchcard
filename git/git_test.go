package git

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInit(t *testing.T) {
	currentDir, _ := os.Getwd()
	testDir := filepath.Join(currentDir, "testDir")
	Init(testDir)
	if !exists(filepath.Join(testDir, ".git")) {
		t.Errorf("After git init, there should be a .git dir.")
	}
	os.RemoveAll(testDir)
}

func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
