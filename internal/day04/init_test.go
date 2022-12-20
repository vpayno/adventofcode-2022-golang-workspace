package day04

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

func TestSection_addRange(t *testing.T) {
	s := section{}
	err := s.addRange("2-4")
	assert.Nil(t, err)

	want1s := 2
	want1e := 4

	got1s := s.start
	got1e := s.end

	assert.Equal(t, want1s, got1s)
	assert.Equal(t, want1e, got1e)
}

func TestPair_addPair(t *testing.T) {
	want := pair{
		elf1: section{
			start: 1,
			end:   2,
		},
		elf2: section{
			start: 4,
			end:   5,
		},
	}

	got := pair{}
	err := got.addPair("1-2,4-5")
	assert.Nil(t, err)

	assert.Equal(t, want.elf1.start, got.elf1.start)
	assert.Equal(t, want.elf1.end, got.elf1.end)

	assert.Equal(t, want.elf2.start, got.elf2.start)
	assert.Equal(t, want.elf2.end, got.elf2.end)
}

func TestPair_isFullyContained(t *testing.T) {
	p := pair{}

	err := p.addPair("3-4,2-5")
	assert.Nil(t, err)

	assert.True(t, p.isFullyContained(1))
	assert.False(t, p.isFullyContained(2))
}
