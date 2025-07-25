package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"syscall"
	"time"
)

const (
	O_EXCL int = syscall.O_EXCL
)

func main() {
	fmt.Println("--CHECK IF FILE EXISTS--")
	path := "file.txts"
	exists, er := FileExists(path)
	fmt.Println("File exists: ", exists)
	fmt.Println("error: ", er)

	fmt.Println("--CREATE NEW FILE USING os.Create()--")
	path = "create.txt"
	err := FileCreate(path)
	if err != nil {
		fmt.Println("FileCreate: Create File Failed")
	}
	fmt.Println("--CREATE NEW FILE USING os.OpenFile()--")
	path = "createifNotExist.txt"
	err = FileIfNotExists(path)
	if err != nil {
		fmt.Println("FileIfNotExists: Create File Failed")
		fmt.Println(err)
	}
	fmt.Println("--Write in file createifNotExist--")
	path = "createifNotExist.txt"
	fmt.Println("Wrote in file createifNotExist.txt", WriteFile(path, []byte("Hello World")))

	fmt.Println("--append createifNotExist--")
	path = "createifNotExist.txt"
	fmt.Println("append in file createifNotExist.txt", AppendFile(path, []byte("Appending Hello World")))

	fmt.Println("--Reading createifNotExist--")
	path = "createifNotExist.txt"
	data, err := ReadFile(path)
	fmt.Println(string(data), err)

	fmt.Println("--FileStat of createifNotExist--")
	path = "createifNotExist.txt"
	fmt.Println(FileStat(path))
	fmt.Println("File Deleted,", err)

	fmt.Println("--Listing files--")
	fmt.Println(ListDir("/Users/anushasg/Documents/Go/FileHandling"))

	fmt.Println("--COPY of createifNotExist--")
	path = "createifNotExist.txt"
	fmt.Println(Copy(path, "new.txt"))
	fmt.Println("File Deleted,", err)

	fmt.Println("--Deleting createifNotExist--")
	path = "createifNotExist.txt"
	err = Delete(path)
	fmt.Println("File Deleted,", err)

	fmt.Println("--Listing files--")
	fmt.Println(ListDir("/Users/anushasg/Documents/Go/FileHandling"))
}

func FileExists(path string) (bool, error) {
	file, err := os.Lstat(path)
	fmt.Println("File contents: ")
	fmt.Println(file) //<nil> even if data exists
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func FileCreate(path string) error {
	file, err := os.Create(path)
	//if already exists, creates a new empty file
	if err != nil {
		return err
	}
	// fmt.Println(file) //&{0xc00006a1e0}
	defer file.Close()
	return nil
}

func FileIfNotExists(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_EXCL|os.O_CREATE, 0666)
	//if already exists, gives an error "open createifNotExist.txt: file exists because of os.O_EXCL"
	// file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	return err
}

func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func AppendFile(path string, data []byte) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, e := file.Write(data)
	if e != nil {
		return e
	}
	return nil
}

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func FileStat(path string) (int64, time.Time, error) {
	data, err := os.Lstat(path)
	if err != nil {
		return -1, time.Time{}, err
	}
	return data.Size(), data.ModTime(), nil

}
func Copy(src, dest string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	fmt.Println("Printing source", source)
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}
	return destination.Sync()
}

func Delete(path string) error {
	return os.Remove(path)
}

func ListDir(path string) ([]fs.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	onlyFiles := make([]fs.DirEntry, 0)
	for _, entry := range entries {
		if !entry.IsDir() {
			onlyFiles = append(onlyFiles, entry)
		}
	}
	return onlyFiles, nil
}
