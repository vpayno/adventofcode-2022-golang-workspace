// Package day01 is the module with the cli logic for the application.
package day01

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
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

	topCalories := getMaxCalories(data)

	showResult(topCalories)

	topThreeCalories := getResultTopThreeCalories(data)

	showResult(topThreeCalories)

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

func showResult(result int) {
	fmt.Printf("%d\n", result)
}
