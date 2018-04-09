package govern

import (
	"strconv"
	"testing"
)

func TestStringConcat(t *testing.T) {
	start := "hello"
	start += " there"

	if start != "hello there" {
		t.Errorf("string concat doesn't behave as you expect, got %v", start)
	}
}

func TestStringConversion(t *testing.T) {
	conversion, _ := strconv.Atoi("39")
	if conversion != 39 {
		t.Errorf("string conversion to int doesn't behave as you expect, got %v", conversion)
	}
}
