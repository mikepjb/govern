package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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
// TODO take argument for pwd or workout git root
// TODO display max only if less than 15 files in path
func Search() {
	var fileList = FilesInTree(".")
	fmt.Printf("%s >>>\n", strconv.Itoa(len(fileList)))
	for _, path := range fileList[0:15] {
		fmt.Printf("%s\n", path)
	}
}
