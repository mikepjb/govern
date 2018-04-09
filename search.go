package govern

import (
	"fmt"
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

// Controlling the terminal - need access to stdin/out
// then we send vt100 commands etc there.

// here we are making space for the fuzzy search interface
func MakeRoom() {
	fmt.Printf("\n>>>\n\n\n\n\n\n\n")
}
