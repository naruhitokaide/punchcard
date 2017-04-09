package main

import (
	"log"

	"github.com/rtzll/punchcard/commands"
)

func main() {
	if err := commands.PunchCardCmd.Execute(); err != nil {
		log.Fatalf("Execuition failed: %v", err)
	}
}
