package day02

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

func TestGetFile(t *testing.T) {
	fileName := "data/day02/day02-input.txt"

	fileRoot, err := os.Getwd()
	assert.Nil(t, err, err)

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
	assert.Nil(t, err, err)

	tmpDir, err := os.MkdirTemp("/tmp", "TestGetFile_NoCW")
	assert.Nil(t, err, err)

	err = os.Chdir(tmpDir)
	assert.Nil(t, err, err)

	err = os.Remove(tmpDir)
	assert.Nil(t, err, err)

	_, err = getFile("data/day02/day02-input.txt")
	assert.NotNil(t, err, "getFile() should have failed here")

	err = os.Chdir(wd)
	var cause string
	if err != nil {
		cause = err.Error()
	}
	assert.Nil(t, err, "failed to return to the original working directory: "+cause)
}

func TestGetScanner(t *testing.T) {
	fileName := "data/day02/day02-input.txt"

	wantScanner, wantErr := getFile(fileName)
	var cause string
	if wantErr != nil {
		cause = wantErr.Error()
	}
	assert.Nil(t, wantErr, "failed to get wantScanner: "+cause)

	want := bufio.NewScanner(wantScanner)

	gotScanner, gotErr := getFile(fileName)
	if gotErr != nil {
		cause = gotErr.Error()
	}
	assert.Nil(t, gotErr, "failed to get gotScanner: "+cause)

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
	wantLen := 2500

	fileName := "data/day02/day02-input.txt"

	file, err := getFile(fileName)
	assert.Nil(t, err, err)

	scanner := getScanner(file)

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

	file, err := getFile(fileName)
	assert.Nil(t, err, err)

	scanner := getScanner(file)

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
