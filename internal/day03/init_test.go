package day03

import (
	"fmt"
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

func TestRucksack_getCollection(t *testing.T) {
	r := rucksack{
		items: "vJrwpWtwJgWrhcsFMMfFFhFp",
	}

	want1 := "vJrwpWtwJgWr"
	want2 := "hcsFMMfFFhFp"

	got1 := r.getCollection(1)
	got2 := r.getCollection(2)

	assert.Equal(t, want1, got1)
	assert.Equal(t, want2, got2)
}

func TestRucksack_getSharedItem(t *testing.T) {
	r := rucksack{
		items: "vJrwpWtwJgWrhcsFMMfFFhFp",
	}

	want := 'p'
	got, err := r.getSharedItem()

	assert.Nil(t, err)
	msg := fmt.Errorf("want char [%c], got char [%c]", want, got)
	assert.Equal(t, want, got, msg)
}

func TestRucksack_getSharedItem_error(t *testing.T) {
	// even, no shared item
	r := rucksack{
		items: "abcdefgh",
	}

	_, err := r.getSharedItem()

	assert.NotNil(t, err)
}

func TestRucksack_getSharedPriority(t *testing.T) {
	r := rucksack{
		items: "vJrwpWtwJgWrhcsFMMfFFhFp",
	}

	want := 16
	got, err := r.getSharedPriority()

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestRucksack_getSharedPriority_error(t *testing.T) {
	// odd length string
	r := rucksack{
		items: "abcdefgh",
	}

	_, err := r.getSharedPriority()

	assert.NotNil(t, err)
}

func TestRucksack_addItems(t *testing.T) {
	want := rucksack{
		items: "vJrwpWtwJgWrhcsFMMfFFhFp",
	}

	got := new(rucksack)
	gotErr := got.addItems("vJrwpWtwJgWrhcsFMMfFFhFp")

	assert.Nil(t, gotErr, "there was an error adding the items to the rucksack")
	assert.Equal(t, want.items, got.items)
}

func TestGetPriority(t *testing.T) {
	char := 'a'
	want := 1
	got := getPriority(char)
	assert.Equal(t, want, got, "wrong priority, "+", for item "+string(char))

	char = 'z'
	want = 26
	got = getPriority(char)
	assert.Equal(t, want, got, "wrong priority, "+", for item "+string(char))

	char = 'A'
	want = 27
	got = getPriority(char)
	assert.Equal(t, want, got, "wrong priority, "+", for item "+string(char))

	char = 'Z'
	want = 52
	got = getPriority(char)
	assert.Equal(t, want, got, "wrong priority, "+", for item "+string(char))
}
