// Package day03 is the module with the cli logic for the application.
package day03

import (
	"bufio"
	"fmt"

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

	groupPrioritySum, err := getGroupPrioritySum(data)
	if err != nil {
		return err
	}

	aocshared.ShowResult(groupPrioritySum)

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

func getGroupPrioritySum(data rucksacks) (int, error) {
	var groupSum int
	group := []string{}
	sacks := [][]string{}

	if len(data) == 0 {
		err := fmt.Errorf("rucksack list is empty: %v", data)
		return 0, err
	}

	if len(data)%3 != 0 {
		err := fmt.Errorf("rucksack list length, %d, isn't evenly divisible by 3", len(data))
		return 0, err
	}

	for i, r := range data {
		if i%3 == 0 && i > 0 {
			sacks = append(sacks, group)
			group = []string{}
		}

		group = append(group, r.items)
	}

	sacks = append(sacks, group)

	for _, group := range sacks {
		// not handling the case where the number of rucksacks isn't evenly divisibleby 3
		g1 := aocshared.SetFromSlice(aocshared.SplitString(group[0]))
		g2 := aocshared.SetFromSlice(aocshared.SplitString(group[1]))
		g3 := aocshared.SetFromSlice(aocshared.SplitString(group[2]))

		c1 := aocshared.SetFromSlice(aocshared.SetIntersect(g1, g2))
		c2 := aocshared.SetFromSlice(aocshared.SetIntersect(g1, g3))
		c3 := aocshared.SetFromSlice(aocshared.SetIntersect(g2, g3))

		c1 = aocshared.SetFromSlice(aocshared.SetIntersect(c1, c2))
		c1 = aocshared.SetFromSlice(aocshared.SetIntersect(c1, c3))

		for item := range c1 {
			priority := getPriority(item)
			groupSum += priority
		}
	}

	return groupSum, nil
}
