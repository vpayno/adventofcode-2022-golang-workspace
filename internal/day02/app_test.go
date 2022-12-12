package day02

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
	conf := Setup("day02")

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
	conf := Setup("day02")

	// Give it bad data.
	conf.inputFileName = "data/day02/day02-input-bad_data.txt"

	err := Run(conf)
	assert.NotNil(t, err, "Run() didn't fail with a bad data error")
}

func TestLoadData(t *testing.T) {
	wantLen := 2500

	fileName := "data/day02/day02-input.txt"

	file, err := aocshared.GetFile(fileName)
	assert.Nil(t, err, err)

	scanner := aocshared.GetScanner(file)

	got, err := loadData(scanner)
	assert.Nil(t, err, err)

	gotLen := len(got)

	assert.Equal(t, wantLen, gotLen)

	for _, r := range got {
		assert.IsType(t, rock, r.them, "their move isn't of type move")
		assert.IsType(t, rock, r.you, "your move isn't of type move")
	}
}

func TestLoadData_badFile(t *testing.T) {
	fileName := "data/day02/day02-input-bad_data.txt"

	file, err := aocshared.GetFile(fileName)
	assert.Nil(t, err, err)

	scanner := aocshared.GetScanner(file)

	_, err = loadData(scanner)
	assert.NotNil(t, err, err)
}

func TestGetTotalScore(t *testing.T) {
	input := rounds{
		round{them: string2move["A"], you: string2move["Y"]},
		round{them: string2move["B"], you: string2move["X"]},
		round{them: string2move["C"], you: string2move["Z"]},
	}

	want := 15

	got := getTotalScore(input)

	assert.Equal(t, want, got)
}
