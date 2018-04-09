package govern

import (
	"testing"
)

func TestTheNumberOfFilesOfProject(t *testing.T) {
	if FilesAvailable() != 5 {
		t.Errorf("Number of files is %v", FilesAvailable())
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
