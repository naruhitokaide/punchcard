package git

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	TESTDIR     = "testDir"
	TESTFILE    = "testFile"
	TESTMESSAGE = "testMessage"
	TESTDATE    = "2005-04-07T22:13:13"
)

func TestInit(t *testing.T) {
	testDir := getTestDir()
	Init(testDir)
	if !exists(filepath.Join(testDir, ".git")) {
		t.Errorf("After git init, there should be a .git dir.")
	}
	os.RemoveAll(testDir)
}

func TestCommit(t *testing.T) {
	testDir := getTestDir()
	Init(testDir)
	testFile := createTestFile()
	Add(testDir, testFile)
	Commit(testDir, TESTMESSAGE, TESTDATE)
	log := filepath.Join(testDir, ".git", "logs", "refs", "heads", "master")

	if !containsMessage(log) {
		t.Errorf("After commiting the commit message should be in the logs.")
	}
	os.RemoveAll(testDir)
}

func containsMessage(logPath string) bool {
	logFile, _ := os.Open(logPath)
	defer logFile.Close()
	logScanner := bufio.NewScanner(logFile)
	for logScanner.Scan() {
		firstLine := logScanner.Text()
		if strings.Contains(firstLine, TESTMESSAGE) {
			return true
		}
		break
	}
	return false
}

func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func createTestFile() string {
	testFile := filepath.Join(getTestDir(), TESTFILE)
	os.Create(testFile)
	return testFile
}

func getTestDir() string {
	currentDir, _ := os.Getwd()
	return filepath.Join(currentDir, TESTDIR)
}
