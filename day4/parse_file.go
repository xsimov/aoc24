package main

import (
	"log"
	"os"
)

func parseFile(path string) string {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(contents)
}
