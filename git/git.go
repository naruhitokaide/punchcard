package git

import (
	"log"
	"os"
	"os/exec"
)

// permission to create a directory and read/write in it
const PERM = 0755

type Git interface {
	GetLocation() string
	Init()
	Add(filename string)
	Commit(msg, date string)
}

type Repo struct {
	Location string
}

// GetLocation returns the location with which the repo has been created
func (git Repo) GetLocation() string {
	return git.Location
}

// Init initializes a git repo in the given path.
// If the path does not already exists, it will be created.
func (git Repo) Init() {
	if err := os.MkdirAll(git.Location, PERM); err != nil {
		log.Fatal(err)
	}
	err := exec.Command("git", "init", git.Location).Run()
	if err != nil {
		log.Fatal(err)
	}
}

// Add wraps the git add call and will change into the path of the git repo
// and add the file given by name.
func (git Repo) Add(filename string) {
	currentDir, _ := os.Getwd()
	os.Chdir(git.Location)
	err := exec.Command("git", "add", filename).Run()
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(currentDir)
}

// Commit will change into the path of the git repo and execute git commit.
// In addition a message and date for the commit are specified.
func (git Repo) Commit(message, date string) {
	currentDir, _ := os.Getwd()
	os.Chdir(git.Location)
	err := exec.Command("git", "commit", "-m", message, "--date", date).Run()
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(currentDir)
}
