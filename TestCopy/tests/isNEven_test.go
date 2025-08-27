package main

//Running tests for PROCESS PACKAGE
import (
	"testing"

	"github.com/anushasgorawar/goLang/processcopy" //Not the literal path. <githubrepo>/package
)

func TestIsNEven(t *testing.T) {
	n := 10
	expected := "Yes"
	recieved := processcopy.IsNEven(n)
	if expected != recieved {
		t.Errorf("Expected: %v, got %v", expected, recieved)
	}
}
