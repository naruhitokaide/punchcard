package git

import (
	"log"
	"os"
	"os/exec"
)

func Init(path string) {
	log.Println("calling init with " + path)
	if err := os.MkdirAll(path, 0777); err != nil {
		log.Fatal(err)
	}
	log.Println("git init " + path)
	exec.Command("git", "init", path).Run()
}

func Add() {
}

func Commit() {
}
