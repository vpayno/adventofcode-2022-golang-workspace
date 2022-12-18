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

// Empty type to avoid ugly struct{}{}
var Empty struct{}

// Set holds a collection of unique items.
type Set map[rune]struct{}

// SetFromSlice creates a set from a slice.
func SetFromSlice(slice []rune) Set {
	set := make(Set)

	for _, element := range slice {
		set[element] = Empty
	}

	return set
}

// SetIntersect finds the intersection of two sets.
func SetIntersect(s1, s2 Set) []rune {
	common := []rune{}

	for key := range s1 {
		_, found := s2[key]
		if found {
			common = append(common, key)
		}
	}

	return common
}

// SplitString converts a string into a slice of runes.
func SplitString(str string) []rune {
	runes := []rune{}

	for _, r := range str {
		runes = append(runes, r)
	}

	return runes
}
