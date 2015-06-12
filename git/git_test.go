package git

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestInit(t *testing.T) {
	fmt.Println("TestInit")
	currentDir, _ := os.Getwd()
	testDir := filepath.Join(currentDir, "testDir")
	fmt.Println("TestDir: " + testDir)
	os.Create(testDir)
	Init(testDir)
	if !exists(filepath.Join(testDir, ".git")) {
		t.Errorf("After git init, there should be a .git dir.")
	}
	os.Remove(testDir)
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
