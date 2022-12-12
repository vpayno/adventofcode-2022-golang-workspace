package day01

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
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
	conf := Setup("day01")

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
	conf := Setup("day01")

	// Give it bad data.
	conf.inputFileName = "data/day01/day01-input-bad_data.txt"

	err := Run(conf)
	assert.NotNil(t, err, "Run() didn't fail with a bad data error")
}

func TestGetFile(t *testing.T) {
	fileName := "data/day01/day01-input.txt"

	fileRoot, err := os.Getwd()
	assert.Nil(t, err, err)

	fileRoot = filepath.Clean(fileRoot + "/../../")
	testFile := fileRoot + "/" + fileName

	wantFile, wantErr := os.Open(testFile)

	gotFile, gotErr := aocshared.GetFile(fileName)

	assert.Nil(t, wantErr, "failed to open wanted file")
	assert.Nil(t, gotErr, "failed to open got file")

	assert.Equal(t, wantFile.Name(), gotFile.Name(), "file names don't match")
}

func TestGetFile_NoCW(t *testing.T) {
	wd, err := os.Getwd()
	assert.Nil(t, err, err)

	tmpDir, err := os.MkdirTemp("/tmp", "TestGetFile_NoCW")
	assert.Nil(t, err, err)

	err = os.Chdir(tmpDir)
	assert.Nil(t, err, err)

	err = os.Remove(tmpDir)
	assert.Nil(t, err, err)

	_, err = aocshared.GetFile("data/day01/day01-input.txt")
	assert.NotNil(t, err, "aocshared.GetFile() should have failed here")

	err = os.Chdir(wd)
	var cause string
	if err != nil {
		cause = err.Error()
	}
	assert.Nil(t, err, "failed to return to the original working directory: "+cause)
}

func TestGetScanner(t *testing.T) {
	fileName := "data/day01/day01-input.txt"

	wantScanner, wantErr := aocshared.GetFile(fileName)
	var cause string
	if wantErr != nil {
		cause = wantErr.Error()
	}
	assert.Nil(t, wantErr, "failed to get wantScanner: "+cause)

	want := bufio.NewScanner(wantScanner)

	gotScanner, gotErr := aocshared.GetFile(fileName)
	if gotErr != nil {
		cause = gotErr.Error()
	}
	assert.Nil(t, gotErr, "failed to get gotScanner: "+cause)

	got := aocshared.GetScanner(gotScanner)

	for {
		if !want.Scan() {
			break
		}
		lineWant := want.Text()

		if !got.Scan() {
			break
		}
		lineGot := got.Text()

		assert.Equal(t, lineWant, lineGot, "lines in "+fileName+" not equal")
	}
}

func TestLoadData(t *testing.T) {
	want := map[string]int{
		"elf1": 51_844,
		"elf2": 44_141,
		"elf3": 28_132,
		"elf4": 49_298,
		"elf5": 26_247,
	}

	fileName := "data/day01/day01-input.txt"

	file, err := aocshared.GetFile(fileName)
	assert.Nil(t, err, err)

	scanner := aocshared.GetScanner(file)

	got, err := loadData(scanner)
	assert.Nil(t, err, err)

	for key, wantValue := range want {
		gotValue, gotFound := got[key]
		assert.True(t, gotFound, "wanted key, "+key+", not found in dictionary")
		assert.Equal(t, wantValue, gotValue, "want/got values don't match for key ["+key+"]")
	}
}

func TestLoadData_badFile(t *testing.T) {
	fileName := "data/day01/day01-input-bad_data.txt"

	file, err := aocshared.GetFile(fileName)
	assert.Nil(t, err, err)

	scanner := aocshared.GetScanner(file)

	_, err = loadData(scanner)
	assert.NotNil(t, err, err)
}

func TestGetMaxCalories(t *testing.T) {
	input := map[string]int{
		"elf1": 2_000,
		"elf2": 8_000,
		"elf3": 5_000,
		"elf4": 1_000,
		"elf5": 3_000,
	}

	want := 8_000

	got := getMaxCalories(input)

	assert.Equal(t, want, got, "highest elf calories don't match")
}

func TestGetTopThreeSum(t *testing.T) {
	input := []int{2_000, 7_000, 3_000, 5_000, 1_000}

	want := 15_000

	got := getTopThreeSum(input)

	assert.Equal(t, want, got, "top three sums don't match")
}

func TestGetResultTopThreeCalories(t *testing.T) {
	input := map[string]int{
		"elf1": 2_000,
		"elf2": 8_000,
		"elf3": 5_000,
		"elf4": 1_000,
		"elf5": 3_000,
	}

	want := 16_000

	got := getResultTopThreeCalories(input)

	assert.Equal(t, want, got, "top three sums don't match")
}

func TestShowResult(t *testing.T) {
	testStdout, writer, err := os.Pipe()
	if err != nil {
		t.Errorf("os.Pipe() err %v; want %v", err, nil)
	}

	osStdout := os.Stdout // keep backup of the real stdout
	os.Stdout = writer

	defer func() {
		// Undo what we changed when this test is done.
		os.Stdout = osStdout
	}()

	want := "1234\n"

	// Run the function who's output we want to capture.
	aocshared.ShowResult(1234)

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("main(); want %q, got %q", want, got)
	}
}
