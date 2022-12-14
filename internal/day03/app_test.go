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

func TestRun_fileError(t *testing.T) {
	// this tests fails because it can't find the file
	conf := Setup("day00")

	err := Run(conf)
	assert.NotNil(t, err, "Run() didn't fail with a can't find file error")
}

func TestRun_badData(t *testing.T) {
	// this tests fails because it can't find the file
	conf := Setup("day03")

	// Give it bad data.
	conf.inputFileName = "data/day03/day03-input-bad_data.txt"

	err := Run(conf)
	assert.NotNil(t, err, "Run() didn't fail with a bad data error")
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
	t.Skip("disabling test for now")

	fileName := "data/day03/day03-input-bad_data2.txt"

	file, err := aocshared.GetFile(fileName)
	assert.Nil(t, err, err)

	scanner := aocshared.GetScanner(file)

	_, err = loadData(scanner)
	assert.NotNil(t, err, err)
}

func TestGetPrioritySum(t *testing.T) {
	rucksacks := []rucksack{}

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
		rucksacks = append(rucksacks, r)
	}

	want := 157

	got, err := getPrioritySum(rucksacks)
	assert.Nil(t, err)

	assert.Equal(t, want, got)
}

func TestGetPrioritySum_error(t *testing.T) {
	rucksacks := []rucksack{}

	r := rucksack{}

	data := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFpx",
	}

	for _, s := range data {
		err := r.addItems(s)
		assert.NotNil(t, err, err)
		rucksacks = append(rucksacks, r)
	}

	_, err := getPrioritySum(rucksacks)
	assert.NotNil(t, err)
}
