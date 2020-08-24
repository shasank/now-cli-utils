package utils

import (
	"fmt"
	"io"
	"os"
)

// Constants
const (
	ver = "v1.4.0"
)

var path = "test.txt"

func main() {
	fmt.Println(GetVersion())
}

// GetVersion ->
func GetVersion() string {
	return ver
}

// VersionCompatibilityCheck ->
func VersionCompatibilityCheck(cliVersion string) {
	if cliVersion != GetVersion() {
		fmt.Println("os.args from utils", os.Args)
		fmt.Println("Incompatible version of core-cli-utils")
		os.Exit(1)
	}
}

// CreateFile ->
func CreateFile() {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("File Created Successfully", path)
}

// WriteFile ->
func WriteFile() {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString("Hello \n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("World \n")
	if isError(err) {
		return
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("File Updated Successfully.")
}

// ReadFile ->
func ReadFile() {
	// Open file for reading.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	fmt.Println("Reading from file.")
	fmt.Println(string(text))
}

// DeleteFile ->
func DeleteFile() {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("File Deleted")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
