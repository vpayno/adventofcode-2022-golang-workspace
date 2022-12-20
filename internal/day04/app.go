// Package day04 is the module with the cli logic for the application.
package day04

import (
	"bufio"

	"github.com/vpayno/adventofcode-2022-golang-workspace/internal/aocshared"
)

// Run is called my the gain function. It's basically the main function of the app.
func Run(conf Config) error {
	file, err := aocshared.GetFile(conf.inputFileName)
	if err != nil {
		return err
	}

	scanner := aocshared.GetScanner(file)

	data, err := loadData(scanner)
	if err != nil {
		return err
	}

	containedSum := getFullyContainedCount(data)

	aocshared.ShowResult(containedSum)

	return nil
}

func loadData(file *bufio.Scanner) (pairs, error) {
	data := pairs{}

	for file.Scan() {
		line := file.Text()

		if line == "" {
			continue
		}

		p := pair{}
		err := p.addPair(line)
		if err != nil {
			return data, err
		}

		data = append(data, p)
	}

	return data, nil
}

func getFullyContainedCount(data pairs) int {
	var count int

	for _, p := range data {
		if p.isFullyContained(1) || p.isFullyContained(2) {
			count++
		}
	}

	return count
}
