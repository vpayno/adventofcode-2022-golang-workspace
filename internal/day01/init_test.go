package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	appName := "testName"

	want := Config{
		appName:       appName,
		inputFileName: "data/" + appName + "/" + appName + "-input.txt",
	}

	got := Setup(appName)

	assert.Equal(t, want.appName, got.appName, "app names aren't equal")
	assert.Equal(t, want.inputFileName, got.inputFileName, "input file names aren't equal")
}
