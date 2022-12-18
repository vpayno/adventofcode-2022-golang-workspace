package day03

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpayno/adventofcode-2022-golang-workspace/internal/aocshared"
)

/*
  Stdout testing code borrowed from Jon Calhoun's FizzBuzz example.
  https://courses.calhoun.io/lessons/les_algo_m01_08
  https://github.com/joncalhoun/algorithmswithgo.com/blob/master/module01/fizz_buzz_test.go
*/

// This is the main test function. This is the gatekeeper of all the tests in the appwc package.
func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestRun(t *testing.T) {
	conf := Setup("day03")

	err := Run(conf)
	var cause string
	if err != nil {
		cause = err.Error()
	}
	assert.Nil(t, err, cause)
}

func TestRun_missingFile(t *testing.T) {
	// this tests fails because it can't find the file
	conf := Setup("day00")

	err := Run(conf)
	assert.NotNil(t, err, "Run() didn't fail with a can't find file error")
}

func TestRun_badData(t *testing.T) {
	// this tests fails because some of the records are invalid
	conf := Setup("day03")

	// Give it bad data.
	conf.inputFileName = "data/day03/day03-input-bad_data.txt"

	err := Run(conf)
	assert.NotNil(t, err, "Run() didn't fail with a bad data error")
}

func TestRun_badData2(t *testing.T) {
	// this tests fails because some of the records are invalid
	conf := Setup("day03")

	// Give it bad data.
	conf.inputFileName = "data/day03/day03-input-uneven_pockets.txt"

	err := Run(conf)
	assert.NotNil(t, err, "Run() didn't fail with an addItems() error")
}

func TestRun_unevenGroups(t *testing.T) {
	// this tests fails because the not all the sack groups have 3 elves
	conf := Setup("day03")

	// Give it bad data.
	conf.inputFileName = "data/day03/day03-input-bad_group_size.txt"

	err := Run(conf)
	assert.NotNil(t, err, "Run() didn't fail with a group size error")
}

func TestRun_noSharedItems(t *testing.T) {
	// this tests fails because there are no shared items
	conf := Setup("day03")

	// Give it bad data.
	conf.inputFileName = "data/day03/day03-input-no_shared_items.txt"

	err := Run(conf)
	assert.NotNil(t, err, "Run() didn't fail with no shared item error")
}

func TestLoadData(t *testing.T) {
	wantLen := 300

	fileName := "data/day03/day03-input.txt"

	file, err := aocshared.GetFile(fileName)
	assert.Nil(t, err, err)

	scanner := aocshared.GetScanner(file)

	got, err := loadData(scanner)
	assert.Nil(t, err, err)
	gotLen := len(got)

	assert.Equal(t, wantLen, gotLen, "read data is the wrong size")
}

func TestLoadData_badFile1(t *testing.T) {
	fileName := "data/day03/day03-input-bad_data1.txt"

	file, err := aocshared.GetFile(fileName)
	assert.Nil(t, err, err)

	scanner := aocshared.GetScanner(file)

	_, err = loadData(scanner)
	assert.NotNil(t, err, err)
}

func TestLoadData_badFile2(t *testing.T) {
	fileName := "data/day03/day03-input-uneven_pockets.txt"

	file, err := aocshared.GetFile(fileName)
	assert.Nil(t, err, err)

	scanner := aocshared.GetScanner(file)

	_, err = loadData(scanner)
	assert.NotNil(t, err, err)
}

func TestGetPrioritySum(t *testing.T) {
	sacks := rucksacks{}

	r := rucksack{}

	data := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	for _, s := range data {
		err := r.addItems(s)
		assert.Nil(t, err, err)
		sacks = append(sacks, r)
	}

	want := 157

	got, err := getPrioritySum(sacks)
	assert.Nil(t, err)

	assert.Equal(t, want, got)
}

func TestGetPrioritySum_error(t *testing.T) {
	sacks := rucksacks{}

	r := rucksack{}

	data := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFpx",
	}

	for _, s := range data {
		err := r.addItems(s)
		assert.NotNil(t, err, err)
		sacks = append(sacks, r)
	}

	_, err := getPrioritySum(sacks)
	assert.NotNil(t, err)
}

func TestGetGroupPrioritySum(t *testing.T) {
	wantSackGroups := 6
	want := 70

	sacks := rucksacks{}

	r := rucksack{}

	data := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	for _, s := range data {
		err := r.addItems(s)
		assert.Nil(t, err, err)
		sacks = append(sacks, r)
	}

	gotSackGroups := len(sacks)
	assert.Equal(t, wantSackGroups, gotSackGroups)

	got, err := getGroupPrioritySum(sacks)
	assert.Nil(t, err)

	assert.Equal(t, want, got)
}

func TestGetGroupPrioritySum_emptyList(t *testing.T) {
	wantSackGroups := 0
	want := 0

	sacks := rucksacks{}

	r := rucksack{}

	data := []string{}

	for _, s := range data {
		err := r.addItems(s)
		assert.Nil(t, err, err)
		sacks = append(sacks, r)
	}

	gotSackGroups := len(sacks)
	assert.Equal(t, wantSackGroups, gotSackGroups)

	got, err := getGroupPrioritySum(sacks)
	assert.NotNil(t, err)

	assert.Equal(t, want, got)
}

func TestGetGroupPrioritySum_notDivisibleByThree(t *testing.T) {
	wantSackGroups := 2
	want := 0

	sacks := rucksacks{}

	r := rucksack{}

	data := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	}

	for _, s := range data {
		err := r.addItems(s)
		assert.Nil(t, err, err)
		sacks = append(sacks, r)
	}

	gotSackGroups := len(sacks)
	assert.Equal(t, wantSackGroups, gotSackGroups)

	got, err := getGroupPrioritySum(sacks)
	assert.NotNil(t, err)

	assert.Equal(t, want, got)
}
