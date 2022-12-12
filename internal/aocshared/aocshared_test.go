package aocshared

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

func TestGetFile(t *testing.T) {
	fileName := "data/day02/day02-input.txt"

	fileRoot, err := os.Getwd()
	assert.Nil(t, err, err)

	fileRoot = filepath.Clean(fileRoot + "/../../")
	testFile := fileRoot + "/" + fileName

	wantFile, wantErr := os.Open(testFile)

	gotFile, gotErr := GetFile(fileName)

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

	_, err = GetFile("data/day02/day02-input.txt")
	assert.NotNil(t, err, "GetFile() should have failed here")

	err = os.Chdir(wd)
	var cause string
	if err != nil {
		cause = err.Error()
	}
	assert.Nil(t, err, "failed to return to the original working directory: "+cause)
}

func TestGetScanner(t *testing.T) {
	fileName := "data/day02/day02-input.txt"

	wantScanner, wantErr := GetFile(fileName)
	var cause string
	if wantErr != nil {
		cause = wantErr.Error()
	}
	assert.Nil(t, wantErr, "failed to get wantScanner: "+cause)

	want := bufio.NewScanner(wantScanner)

	gotScanner, gotErr := GetFile(fileName)
	if gotErr != nil {
		cause = gotErr.Error()
	}
	assert.Nil(t, gotErr, "failed to get gotScanner: "+cause)

	got := GetScanner(gotScanner)

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
	ShowResult(1234)

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
