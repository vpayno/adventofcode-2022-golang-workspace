// Package aocshared is the module with the cli logic for the application.
package aocshared

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// GetFile takes a file name and returns a file handle.
func GetFile(fileName string) (*os.File, error) {
	fileRoot, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fileRoot = filepath.Clean(fileRoot + "/../../")
	fileName = filepath.Clean(fileRoot + "/" + fileName)

	// strings.HasPrefix() is pointless since we're generating the full path.

	// https://securego.io/docs/rules/g304.html - this gosec check seems to want
	// panic() calls
	file, err := os.Open(fileName) // #nosec

	return file, err
}

// GetScanner takes a file handle and returns a bufio scanner.
func GetScanner(file *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return scanner
}

// ShowResult prints the passed integer.
func ShowResult(result int) {
	fmt.Printf("%d\n", result)
}
