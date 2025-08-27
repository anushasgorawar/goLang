package main

//Running tests for PROCESS PACKAGE
import (
	"testing"

	process "github.com/anushasgorawar/goLang/Process"
)

func TestIsNEven(t *testing.T) {
	n := 10
	expected := "Yess"
	recieved := process.IsNEven(n)
	if expected != recieved {
		t.Errorf("Expected: %v, got %v", expected, recieved)
	}
}

// (base) Anushas-MacBook-Air:Test anushasg$ go test .
// ok      github.com/anushasgorawar/goLang        0.661s
// (base) Anushas-MacBook-Air:Test anushasg$ go test .
// --- FAIL: TestIsEven (0.00s)
//     even_test.go:10: Expected: Yess, got Yes
// FAIL
// FAIL    github.com/anushasgorawar/goLang        0.637s
// FAIL
