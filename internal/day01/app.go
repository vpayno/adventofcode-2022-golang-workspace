// Package day01 is the module with the cli logic for the application.
package day01

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"

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

	topCalories := getMaxCalories(data)

	aocshared.ShowResult(topCalories)

	topThreeCalories := getResultTopThreeCalories(data)

	aocshared.ShowResult(topThreeCalories)

	return nil
}

func loadData(file *bufio.Scanner) (map[string]int, error) {
	var err error

	data := map[string]int{}

	var elfCounter = 1
	var calories int

	for file.Scan() {
		line := file.Text()

		elfName := fmt.Sprintf("%s%d", "elf", elfCounter)

		if line == "" {
			elfCounter++
			data[elfName] = calories
			calories = 0

			continue
		}

		calorie, err := strconv.Atoi(line)
		if err != nil {
			return map[string]int{}, err
		}

		calories += calorie
	}

	return data, err
}

func getMaxCalories(data map[string]int) int {
	var result int

	for _, calories := range data {
		if calories > result {
			result = calories
		}
	}

	return result
}

func getTopThreeSum(calories []int) int {
	var sum int

	var start int

	if len(calories) > 3 {
		start += len(calories) - 3
	}

	sort.Ints(calories)

	for _, calorie := range calories[start:] {
		sum += calorie
	}

	return sum
}

func getResultTopThreeCalories(data map[string]int) int {
	totals := []int{}

	for _, calories := range data {
		totals = append(totals, calories)
	}

	result := getTopThreeSum(totals)

	return result
}
