package git

import (
	"log"
	"os"
	"os/exec"
)

// permission to create a directory and read/write in it
const PERM = 0755

type Git interface {
	Init()
	Add(filename string)
	Commit(msg, date string)
}

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

// Add wraps the git add call and will change into the location of the git repo
// add the file given by name and change back to the previous directory.
func (git Repo) Add(filename string) {
	currentDir, _ := os.Getwd()
	os.Chdir(git.Location)
	err := exec.Command("git", "add", filename).Run()
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(currentDir)
}

// Commit will change into the location of the git repo and execute git commit
// with a message and date and change back to the previous directory.
func (git Repo) Commit(message, date string) {
	currentDir, _ := os.Getwd()
	os.Chdir(git.Location)
	err := exec.Command("git", "commit", "-m", message, "--date", date).Run()
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(currentDir)
}
