package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FilesAvailable() int {
	return 5
}

var buffer = ""

func InsertCharacter(input string) {
	buffer += input
}

func GetBuffer() string {
	return buffer
}

func FilesInTree(root string) []string {
	var files []string

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !strings.Contains(path, ".git") && !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files
}

// Controlling the terminal - need access to stdin/out
// then we send vt100 commands etc there.

// here we are making space for the fuzzy search interface
func MakeRoom() {
	fmt.Printf(">>>\n\n\n\n\n\n\n\n\n\n\n")
}
