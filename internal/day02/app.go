// Package day02 is the module with the cli logic for the application.
package day02

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Run is called my the gain function. It's basically the main function of the app.
func Run(conf Config) error {
	file, err := getFile(conf.inputFileName)
	if err != nil {
		return err
	}

	scanner := getScanner(file)

	data, err := loadData(scanner)
	if err != nil {
		return err
	}

	totalScore := getTotalScore(data)

	showResult(totalScore)

	return nil
}

func getFile(fileName string) (*os.File, error) {
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

func getScanner(file *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return scanner
}

func loadData(file *bufio.Scanner) (rounds, error) {
	var err error

	data := rounds{}

	for file.Scan() {
		line := file.Text()

		if line == "" {
			continue
		}

		s := strings.Fields(line)

		if len(s) != 2 {
			err := fmt.Errorf("wrong number of records: has %d, need %d", len(s), 2)
			return []round{}, err
		}

		r := round{
			them: string2move[s[them]],
			you:  string2move[s[you]],
		}

		data = append(data, r)
	}

	return data, err
}

func getTotalScore(data rounds) int {
	var result int

	for _, r := range data {
		result += r.score()
	}

	return result
}

func showResult(result int) {
	fmt.Printf("%d\n", result)
}
