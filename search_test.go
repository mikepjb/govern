package main

import (
	"testing"
)

func sliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TestTheNumberOfFilesOfProject(t *testing.T) {
	if FilesAvailable() != 5 {
		t.Errorf("Number of files is %v", FilesAvailable())
	}
}

func TestListFilesInTree(t *testing.T) {
	var fileList = FilesInTree(".")

	if !sliceContains(fileList, "govern.go") {
		t.Errorf("file list does not contain govern.go")
	}

	if !sliceContains(fileList, "search_test.go") {
		t.Errorf("file list does not contain search_test.go")
	}
}

func TestInsertCharacter(t *testing.T) {
	if GetBuffer() != "" {
		t.Errorf("Buffer is not blank at the start")
	}
	InsertCharacter("r")
	InsertCharacter("e")
	InsertCharacter("a")
	InsertCharacter("d")
	if GetBuffer() != "read" {
		t.Errorf("Buffer is not \"read\"")
	}
}
