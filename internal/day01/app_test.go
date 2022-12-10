package day01

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.Nil(t, err, "Run() returned an error")
}

func TestRun_fileError(t *testing.T) {
	// this tests fails because it can't find the file
	conf := Setup("day00")

	err := Run(conf)
	assert.NotNil(t, err, "Run() didn't fail")
}

func TestGetFile(t *testing.T) {
	fileName := "data/day01/day01-input.txt"

	fileRoot, err := os.Getwd()
	assert.Nil(t, err, "failed to get CWD")

	fileRoot = filepath.Clean(fileRoot + "/../../")
	testFile := fileRoot + "/" + fileName

	wantFile, wantErr := os.Open(testFile)

	gotFile, gotErr := getFile(fileName)

	assert.Nil(t, wantErr, "failed to open wanted file")
	assert.Nil(t, gotErr, "failed to open got file")

	assert.Equal(t, wantFile.Name(), gotFile.Name(), "file names don't match")
}

func TestGetFile_NoCW(t *testing.T) {
	wd, err := os.Getwd()
	assert.Nil(t, err, "failed to get CWD, should always work here")

	tmpDir, err := os.MkdirTemp("/tmp", "TestGetFile_NoCW")
	assert.Nil(t, err, "failed to create temporary directory")

	err = os.Chdir(tmpDir)
	assert.Nil(t, err, "failed to cd to temporary directory")

	err = os.Remove(tmpDir)
	assert.Nil(t, err, "failed to remove temporary directory")

	_, err = getFile("data/day01/day01-input.txt")
	assert.NotNil(t, err, "getFile() should have failed here")

	err = os.Chdir(wd)
	assert.Nil(t, err, "failed to return to the orginal working directory")
}

func TestGetScanner(t *testing.T) {
	fileName := "data/day01/day01-input.txt"

	wantScanner, wantErr := getFile(fileName)
	assert.Nil(t, wantErr, "failed to get wantScanner")

	want := bufio.NewScanner(wantScanner)

	gotScanner, gotErr := getFile(fileName)
	assert.Nil(t, gotErr, "failed to get gotScanner")

	got := getScanner(gotScanner)

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

	file, err := getFile(fileName)
	assert.Nil(t, err, "failed to get file")

	scanner := getScanner(file)

	got, err := loadData(scanner)
	assert.Nil(t, err, "failed to load file")

	for key, wantValue := range want {
		gotValue, gotFound := got[key]
		assert.True(t, gotFound, "wanted key, "+key+", not found in dictionary")
		assert.Equal(t, wantValue, gotValue, "want/got values don't match for key ["+key+"]")
	}
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
	showResult(1234)

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
