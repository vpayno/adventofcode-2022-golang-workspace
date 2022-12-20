// Package day04 is the module with the cli logic for the cmd application.
package day04

import (
	"fmt"
	"strconv"
	"strings"
)

// Config holds the application's configuration.
type Config struct {
	appName       string
	inputFileName string
}

// Section stores the start and end range of an Elf's section.
type section struct {
	start int
	end   int
}

func (s *section) addRange(r string) error {
	data := strings.Split(r, "-")

	if len(data) != 2 {
		return fmt.Errorf("range did not contain start and end values (%v)", data)
	}

	var err error

	start, err := strconv.Atoi(data[0])
	if err != nil {
		return err
	}

	end, err := strconv.Atoi(data[1])
	if err != nil {
		return err
	}

	if start < 1 {
		return fmt.Errorf("range start, %d, is less than 1", start)
	}

	if end > 99 {
		return fmt.Errorf("range end, %d, is greather than 99", end)
	}

	s.start = start
	s.end = end

	return nil
}

// Pair holds 2 elves sections.
type pair struct {
	elf1 section
	elf2 section
}

func (p *pair) addPair(input string) error {
	data := strings.Split(input, ",")

	if len(data) != 2 {
		return fmt.Errorf("input list doesn't contain two entries (%v)", data)
	}

	err := p.elf1.addRange(data[0])
	if err != nil {
		return err
	}

	err = p.elf2.addRange(data[1])
	if err != nil {
		return err
	}

	return nil
}

func (p *pair) isFullyContained(elfNo int) bool {
	if elfNo == 1 {
		return p.elf1.start >= p.elf2.start && p.elf1.end <= p.elf2.end
	}

	return p.elf2.start >= p.elf1.start && p.elf2.end <= p.elf1.end
}

func (p *pair) isPartiallyContained(elfNo int) bool {
	if elfNo == 1 {
		return (p.elf2.start <= p.elf1.start && p.elf1.start <= p.elf2.end) || (p.elf2.start <= p.elf1.end && p.elf1.end <= p.elf2.end)
	}

	return (p.elf1.start <= p.elf2.start && p.elf2.start <= p.elf1.end) || (p.elf1.start <= p.elf2.end && p.elf2.end <= p.elf1.end)
}

type pairs []pair

// Setup creates the applications configuration object.
func Setup(appName string) Config {

	conf := Config{
		appName:       appName,
		inputFileName: "data/" + appName + "/" + appName + "-input.txt",
	}

	return conf
}
