package main

import (
	"io"
	"log"
	"os"
)

func parseFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return string(bytes)
}
