package main

import (
	"github.com/0xfoo/punchcard/commands"
	"log"
)

func main() {
	if err := commands.PunchCardCmd.Execute(); err != nil {
		log.Fatalf("Execuition failed: %v", err)
	}
}
