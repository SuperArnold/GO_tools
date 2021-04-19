package main

import (
	"log"

	"github.com/SuperArnold/GO_tools/2.Word_format/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
