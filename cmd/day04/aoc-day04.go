// Package main is the cli module for the application.
package main

import (
	"fmt"

	app "github.com/vpayno/adventofcode-2022-golang-workspace/internal/day04"
)

var challengeName = "day04"

func main() {
	// Create the configuration object.
	conf := app.Setup(challengeName)

	// Run the main application.
	err := app.Run(conf)
	if err != nil {
		fmt.Println("Encountered error while running app.Run()")
		fmt.Println()
		fmt.Println(err)

		// Don't panic
		return
	}
}
