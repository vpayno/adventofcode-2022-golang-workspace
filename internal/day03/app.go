// Package day03 is the module with the cli logic for the application.
package day03

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

	prioritySum, err := getPrioritySum(data)
	if err != nil {
		return err
	}

	aocshared.ShowResult(prioritySum)

	return nil
}

func loadData(file *bufio.Scanner) (rucksacks, error) {
	data := rucksacks{}

	for file.Scan() {
		line := file.Text()

		if line == "" {
			continue
		}

		r := rucksack{}
		err := r.addItems(line)
		if err != nil {
			return data, err
		}

		data = append(data, r)
	}

	return data, nil
}

func getPrioritySum(data rucksacks) (int, error) {
	var result int

	for _, r := range data {
		priority, err := r.getSharedPriority()
		if err != nil {
			return 0, err
		}

		result += priority
	}

	return result, nil
}
