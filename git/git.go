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

func Add() {
}

func Commit() {
}
