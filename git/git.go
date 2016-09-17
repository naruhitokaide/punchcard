package git

import (
	"log"
	"os"
	"os/exec"
)

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
	if err := os.MkdirAll(git.Location, 0755); err != nil {
		log.Fatal(err)
	}
	if err := exec.Command("git", "init", git.Location).Run(); err != nil {
		log.Fatal(err)
	}
}

// Add wraps the git add call and will change into the location of the git repo
// add the file given by name and change back to the previous directory.
func (git Repo) Add(filename string) {
	cmd := exec.Command("git", "add", filename)
	cmd.Dir = git.Location
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

// Commit will change into the location of the git repo and execute git commit
// with a message and date and change back to the previous directory.
func (git Repo) Commit(message, date string) {
	cmd := exec.Command("git", "commit", "-m", message, "--date", date)
	cmd.Dir = git.Location
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
