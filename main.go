package main

import (
	"log"

	"github.com/0xfoo/punchcard/commands"
)

func main() {
	if err := commands.PunchCardCmd.Execute(); err != nil {
		log.Fatalf("Execuition failed: %v", err)
	}
}
