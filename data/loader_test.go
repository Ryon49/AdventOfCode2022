package data

import (
	"testing"
)

func TestLoad(t *testing.T) {
	lines, err := Load(0)
	if err != nil {
		t.Error(err)
	}
	if lines[0] != "This is a test" || lines[1] != "123" {
		t.Error("Wrong content")
	}
}
