package day01

import (
	"os"
	"testing"
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

/*
// Use this to put modules, functions in testing mode.
func setupTestEnv() {
}

// Use this to undo things you did in setupTestEnv()
func teardownTestEnv() {
}

func TestRunAppFlagByteCharWordAndLineLongOpts(t *testing.T) {
	setupTestEnv()
	defer teardownTestEnv()

	want := config{
		modes: map[string]bool{
			"byte": true,
			"char": true,
			"line": true,
			"word": true,
		},
	}

	os.Args = []string{"test", "--bytes", "--chars", "--lines", "--words"}
	got, _ := setup()

	if got.modes["byte"] != want.modes["byte"] {
		t.Errorf("setup flags --chars --bytes --lines --words: want %v, got %v", want.modes["byte"], got.modes["byte"])
	}

	if got.modes["char"] != want.modes["char"] {
		t.Errorf("setup flags --chars --bytes --lines --words: want %v, got %v", want.modes["char"], got.modes["char"])
	}

	if got.modes["line"] != want.modes["line"] {
		t.Errorf("setup flags --chars --bytes --lines --words: want %v, got %v", want.modes["line"], got.modes["line"])
	}

	if got.modes["word"] != want.modes["word"] {
		t.Errorf("setup flags --chars --bytes --lines --words: want %v, got %v", want.modes["word"], got.modes["word"])
	}
}
*/
