package git

import (
	"log"
	"os"
	"os/exec"
)

// PERM is the permission to create a directory/file for which the user
// the user is allowed to read from/write to it.
const PERM = 0755

// Git is a simple interface the git command line tool.
type Git interface {
	Init()
	Add(filename string)
	Commit(msg, date string)
}

// Repo represents a git repository by a given location.
type Repo struct {
	Location string
}

// Init initializes a git repo in its location.
// If the location does not already exists, it will be created.
func (git Repo) Init() {
	if err := os.MkdirAll(git.Location, PERM); err != nil {
		log.Fatal(err)
	}
	err := exec.Command("git", "init", git.Location).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func getCwd() string {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return currentWorkingDirectory
}

func changeDir(newDir string) {
	if err := os.Chdir(newDir); err != nil {
		log.Fatal(err)
	}
}

// Add wraps the git add call and will change into the location of the git repo
// add the file given by name and change back to the previous directory.
func (git Repo) Add(filename string) {
	currentDir := getCwd()
	changeDir(git.Location)
	if err := exec.Command("git", "add", filename).Run(); err != nil {
		log.Fatal(err)
	}
	changeDir(currentDir)
}

// Commit will change into the location of the git repo and execute git commit
// with a message and date and change back to the previous directory.
func (git Repo) Commit(message, date string) {
	currentDir := getCwd()
	changeDir(git.Location)
	if err := exec.Command("git", "commit", "-m", message, "--date", date).Run(); err != nil {
		log.Fatal(err)
	}
	changeDir(currentDir)
}
