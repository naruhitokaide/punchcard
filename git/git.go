package git

import (
	"log"
	"os"
	"os/exec"
)

func Init(path string) {
	if err := os.MkdirAll(path, os.ModeDir); err != nil {
		log.Fatal(err)
	}
	err := exec.Command("git", "init", path).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Add(path, filename string) {
	os.Chdir(path)
	err := exec.Command("git", "add", filename).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Commit(path, message, date string) {
	os.Chdir(path)
	err := exec.Command("git", "commit", "-m", message, "--date", date).Run()
	if err != nil {
		log.Fatal(err)
	}
}
