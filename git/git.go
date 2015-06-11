package git

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func Init(path string) {
	if !filepath.IsAbs(path) {
		fmt.Println("Path must be absolute. Given path was: " + path)
		return
	}
	cd := exec.Command("cd", path)
	init := exec.Command("git", "init")

	cd.Run()
	init.Run()
}

func Add() {
}

func Commit() {
}
