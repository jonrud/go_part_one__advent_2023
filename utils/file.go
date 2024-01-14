package utils

import (
	"log"
	"os"
	"strings"
)

// ReadFile Not a very useful wrapper, but  I wanted to have one
func ReadFile(path string) string {
	log.SetPrefix("utils/file: ")

	trimmedPath := strings.TrimSpace(path)
	if "" == trimmedPath {
		log.Fatal("path to file cannot be empty")
	}
	content, err := os.ReadFile(path)

	if nil != err {
		log.Fatalf("failed reading the file: %v\n", err)
	}

	return convertToUnixLineEndings(string(content))
}

func convertToUnixLineEndings(stringToConvert string) string {
	return strings.ReplaceAll(stringToConvert, "\r\n", "\n")
}
