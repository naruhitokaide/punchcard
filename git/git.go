package git

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func Init(path string) {
	log.Println("Entering git init with path: " + path)
	if !filepath.IsAbs(path) {
		log.Fatal("Path must be absolute. Given path was: " + path)
	}
	os.Chdir(path)
	exec.Command("git", "init").Run()
}

func Add() {
}

func Commit() {
}
