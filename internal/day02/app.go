// Package day02 is the module with the cli logic for the application.
package day02

import (
	"bufio"
	"fmt"
	"strings"

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

	totalScore := getTotalScore(data)

	aocshared.ShowResult(totalScore)

	return nil
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
