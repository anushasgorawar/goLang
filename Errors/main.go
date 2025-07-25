package main

import (
	"fmt"
	"os"
	"strings"
)

type NotTextFile struct {
	Message string
}

// type error interface{
// 	Error() string
// }

func (f *NotTextFile) Error() string {
	return f.Message
}

func IsATextFile(path string) (string, error) {
	if !strings.HasSuffix(path, ".txt") {
		return path, &NotTextFile{"Not a Text File"}
	}
	_, err := os.Open(path)
	if err != nil {
		return path, &NotTextFile{"File does not exist"}
	}
	return path, nil
}
func main() {
	path1 := "mod.go"
	path2 := "mood.txt"
	path3 := "mod.txt"

	file, err := IsATextFile(path1)
	fmt.Println(file, err)
	file, err = IsATextFile(path2)
	fmt.Println(file, err)
	nfile, nerr := IsATextFile(path3)
	fmt.Println(nfile, nerr)
}
