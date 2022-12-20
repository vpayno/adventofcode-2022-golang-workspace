package day04

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
	conf := Setup("day04")

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

func TestRun_unevenPairs(t *testing.T) {
	// this tests fails because the not all the sack groups have 3 elves
	conf := Setup("day04")

	// Give it bad data.
	conf.inputFileName = "data/day04/day04-input-bad_pair_size.txt"

	err := Run(conf)
	assert.NotNil(t, err, "Run() didn't fail with a pair size error")
}

func TestRun_notANumberStart(t *testing.T) {
	// this tests fails because there are no shared items
	conf := Setup("day04")

	// Give it bad data.
	conf.inputFileName = "data/day04/day04-input-not_a_number1.txt"

	err := Run(conf)
	assert.NotNil(t, err, "Run() didn't fail with a string to int conversion error")
}

func TestRun_notANumberEnd(t *testing.T) {
	// this tests fails because there are no shared items
	conf := Setup("day04")

	// Give it bad data.
	conf.inputFileName = "data/day04/day04-input-not_a_number2.txt"

	err := Run(conf)
	assert.NotNil(t, err, "Run() didn't fail with a string to int conversion error")
}

func TestRun_incompleteRange(t *testing.T) {
	// this tests fails because there are no shared items
	conf := Setup("day04")

	// Give it bad data.
	conf.inputFileName = "data/day04/day04-input-incomplete_range.txt"

	err := Run(conf)
	assert.NotNil(t, err, "Run() didn't fail with an incomplete range error")
}

func TestRun_badStartRange(t *testing.T) {
	// this tests fails because there are no shared items
	conf := Setup("day04")

	// Give it bad data.
	conf.inputFileName = "data/day04/day04-input-bad_range_start.txt"

	err := Run(conf)
	assert.NotNil(t, err, "Run() didn't fail with a bad range start value error")
}

func TestRun_badEndRange(t *testing.T) {
	// this tests fails because there are no shared items
	conf := Setup("day04")

	// Give it bad data.
	conf.inputFileName = "data/day04/day04-input-bad_range_end.txt"

	err := Run(conf)
	assert.NotNil(t, err, "Run() didn't fail with a bad range start value error")
}

func TestLoadData(t *testing.T) {
	wantLen := 1_000

	fileName := "data/day04/day04-input.txt"

	file, err := aocshared.GetFile(fileName)
	assert.Nil(t, err, err)

	scanner := aocshared.GetScanner(file)

	got, err := loadData(scanner)
	assert.Nil(t, err, err)
	gotLen := len(got)

	assert.Equal(t, wantLen, gotLen, "read data is the wrong size")
}

func TestLoadData_blankLine(t *testing.T) {
	wantLen := 4

	fileName := "data/day04/day04-input-blank_line.txt"

	file, err := aocshared.GetFile(fileName)
	assert.Nil(t, err, err)

	scanner := aocshared.GetScanner(file)

	got, err := loadData(scanner)
	assert.Nil(t, err, err)
	gotLen := len(got)

	assert.Equal(t, wantLen, gotLen, "read data is the wrong size")
}

func TestLoadData_badPairSize(t *testing.T) {
	fileName := "data/day04/day04-input-bad_pair_size.txt"

	file, err := aocshared.GetFile(fileName)
	assert.Nil(t, err, err)

	scanner := aocshared.GetScanner(file)

	_, err = loadData(scanner)
	assert.NotNil(t, err, err)
}

func TestLoadData_notANumberStart(t *testing.T) {
	fileName := "data/day04/day04-input-not_a_number1.txt"

	file, err := aocshared.GetFile(fileName)
	assert.Nil(t, err, err)

	scanner := aocshared.GetScanner(file)

	_, err = loadData(scanner)
	assert.NotNil(t, err, err)
}

func TestLoadData_notANumberEnd(t *testing.T) {
	fileName := "data/day04/day04-input-not_a_number2.txt"

	file, err := aocshared.GetFile(fileName)
	assert.Nil(t, err, err)

	scanner := aocshared.GetScanner(file)

	_, err = loadData(scanner)
	assert.NotNil(t, err, err)
}

func TestLoadData_incompleteRange(t *testing.T) {
	fileName := "data/day04/day04-input-incomplete_range.txt"

	file, err := aocshared.GetFile(fileName)
	assert.Nil(t, err, err)

	scanner := aocshared.GetScanner(file)

	_, err = loadData(scanner)
	assert.NotNil(t, err, err)
}

func TestLoadData_badRangeStart(t *testing.T) {
	fileName := "data/day04/day04-input-bad_range_start.txt"

	file, err := aocshared.GetFile(fileName)
	assert.Nil(t, err, err)

	scanner := aocshared.GetScanner(file)

	_, err = loadData(scanner)
	assert.NotNil(t, err, err)
}

func TestLoadData_badRangeEnd(t *testing.T) {
	fileName := "data/day04/day04-input-bad_range_end.txt"

	file, err := aocshared.GetFile(fileName)
	assert.Nil(t, err, err)

	scanner := aocshared.GetScanner(file)

	_, err = loadData(scanner)
	assert.NotNil(t, err, err)
}

func TestGetFullyContainedCount(t *testing.T) {
	groups := pairs{}

	p := pair{}

	data := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}

	for _, line := range data {
		err := p.addPair(line)
		assert.Nil(t, err, err)
		groups = append(groups, p)
	}

	want := 2

	got := getFullyContainedCount(groups)

	assert.Equal(t, want, got)
}
