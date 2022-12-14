// Package day03 is the module with the cli logic for the cmd application.
package day03

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// Config holds the application's configuration.
type Config struct {
	appName       string
	inputFileName string
}

// Rucksack holds an Elf's items.
type rucksack struct {
	items string
}

func (r *rucksack) addItems(items string) error {
	if len(items)%2 != 0 {
		return fmt.Errorf("items list isn't evenly divisible (len=%d)", len(items))
	}

	r.items = items

	return nil
}

func (r *rucksack) getCollection(num int) string {

	if num == 1 {
		return r.items[0 : len(r.items)/2]
	}

	return r.items[len(r.items)/2:]
}

// lot's of room or optimizations
func (r *rucksack) getSharedItem() (rune, error) {
	collection1 := r.getCollection(1)
	collection2 := r.getCollection(2)

	for _, char := range collection1 {
		if strings.ContainsRune(collection2, char) {
			return char, nil
		}
	}

	return rune(0), errors.New("rucksack compartments doesn't have a shared item")
}

func (r *rucksack) getSharedPriority() (int, error) {
	shared, err := r.getSharedItem()
	if err != nil {
		return 0, err
	}

	if unicode.IsLower(shared) {
		return int(shared) - 96, nil
	}

	return int(shared) - 38, nil
}

type rucksacks []rucksack

// Setup creates the applications configuration object.
func Setup(appName string) Config {

	conf := Config{
		appName:       appName,
		inputFileName: "data/" + appName + "/" + appName + "-input.txt",
	}

	return conf
}
