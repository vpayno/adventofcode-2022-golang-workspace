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

func TestGetFile(t *testing.T) {
	fileName := "data/day01/day01-input.txt"

	fileRoot, err := os.Getwd()
	assert.Nil(t, err)

	fileRoot = filepath.Clean(fileRoot + "/../../")
	testFile := fileRoot + "/" + fileName

	wantFile, wantErr := os.Open(testFile)

	gotFile, gotErr := getFile(fileName)

	assert.Nil(t, wantErr)
	assert.Nil(t, gotErr)

	assert.Equal(t, wantFile.Name(), gotFile.Name(), "file names don't match")
}

func TestGetFile_NoCW(t *testing.T) {
	wd, err := os.Getwd()
	assert.Nil(t, err)

	tmpDir, err := os.MkdirTemp("/tmp", "TestGetFile_NoCW")
	assert.Nil(t, err)

	err = os.Chdir(tmpDir)
	assert.Nil(t, err)

	err = os.Remove(tmpDir)
	assert.Nil(t, err)

	_, err = getFile("data/day01/day01-input.txt")
	assert.NotNil(t, err)

	err = os.Chdir(wd)
	assert.Nil(t, err)
}

func TestGetScanner(t *testing.T) {
	fileName := "data/day01/day01-input.txt"

	wantScanner, wantErr := getFile(fileName)
	want := bufio.NewScanner(wantScanner)

	assert.Nil(t, wantErr)

	gotScanner, gotErr := getFile(fileName)
	got := getScanner(gotScanner)

	assert.Nil(t, gotErr)

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
}

func TestGetMaxCalories(t *testing.T) {
	input := map[string]int{
		"elf1": 2_000,
		"elf2": 8_000,
		"elf3": 5_000,
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
