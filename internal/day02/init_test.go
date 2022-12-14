package day02

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

// Seems like a silly test but any changes to these valuse will break the rest of the logic.
func TestPronoun_values(t *testing.T) {
	assert.Equal(t, int(them), 0)
	assert.Equal(t, int(goal), 1)
}

// Seems like a silly test but any changes to these valuse will break the rest of the logic.
func TestOutcomes_values(t *testing.T) {
	assert.Equal(t, int(loose), 0)
	assert.Equal(t, int(draw), 3)
	assert.Equal(t, int(win), 6)
}

// Seems like a silly test but any changes to these valuse will break the rest of the logic.
func TestOutcomes_strings(t *testing.T) {
	assert.Equal(t, outcomeNames[loose], "loose")
	assert.Equal(t, outcomeNames[draw], "draw")
	assert.Equal(t, outcomeNames[win], "win")
}

// Seems like a silly test but any changes to these valuse will break the rest of the logic.
func TestMoves_values(t *testing.T) {
	assert.Equal(t, int(rock), 1)
	assert.Equal(t, int(paper), 2)
	assert.Equal(t, int(scissors), 3)
}

// Seems like a silly test but any changes to these valuse will break the rest of the logic.
func TestMoves_strings(t *testing.T) {
	assert.Equal(t, moveNames[rock], "rock")
	assert.Equal(t, moveNames[paper], "paper")
	assert.Equal(t, moveNames[scissors], "scissors")
}

// Seems like a silly test but any changes to these valuse will break the rest of the logic.
func TestString2Move(t *testing.T) {
	assert.Equal(t, int(string2move["A"]), int(rock))
	assert.Equal(t, int(string2move["B"]), int(paper))
	assert.Equal(t, int(string2move["C"]), int(scissors))
}

// Seems like a silly test but any changes to these valuse will break the rest of the logic.
func TestString2Outcome(t *testing.T) {
	assert.Equal(t, int(string2outcome["X"]), int(loose))
	assert.Equal(t, int(string2outcome["Y"]), int(draw))
	assert.Equal(t, int(string2outcome["Z"]), int(win))
}

func TestRound_yourMove(t *testing.T) {
	r := round{
		them: rock,
		goal: win,
	}
	assert.Equal(t, int(paper), int(r.yourMove()), "paper vs rock != win")

	r = round{
		them: scissors,
		goal: win,
	}
	assert.Equal(t, int(rock), int(r.yourMove()), "rock vs scissors != win")

	r = round{
		them: paper,
		goal: win,
	}
	assert.Equal(t, int(scissors), int(r.yourMove()), "scissors vs paper != win")

	r = round{
		them: rock,
		goal: loose,
	}
	assert.Equal(t, int(scissors), int(r.yourMove()), "scissors vs rock != loose")

	r = round{
		them: scissors,
		goal: loose,
	}
	assert.Equal(t, int(paper), int(r.yourMove()), "paper vs scissors != loose")

	r = round{
		them: paper,
		goal: loose,
	}
	assert.Equal(t, int(rock), int(r.yourMove()), "rock vs paper != loose")

	r = round{
		them: rock,
		goal: draw,
	}
	assert.Equal(t, int(r.them), int(r.yourMove()), "scissors vs rock != draw")

	r = round{
		them: scissors,
		goal: draw,
	}
	assert.Equal(t, int(r.them), int(r.yourMove()), "scissors vs scissors != draw")

	r = round{
		them: paper,
		goal: draw,
	}
	assert.Equal(t, int(r.them), int(r.yourMove()), "paper vs paper != draw")
}

func TestRound_score(t *testing.T) {
	r := round{
		them: paper,
		goal: loose,
	}
	assert.Equal(t, r.score(), int(loose)+int(rock), "loose+rock != 1")

	r = round{
		them: scissors,
		goal: loose,
	}
	assert.Equal(t, r.score(), int(loose)+int(paper), "loose+paper != 2")

	r = round{
		them: rock,
		goal: loose,
	}
	assert.Equal(t, r.score(), int(loose)+int(scissors), "loose+scissors != 3")

	r = round{
		them: rock,
		goal: draw,
	}
	assert.Equal(t, r.score(), int(draw)+int(rock), "draw+rock != 4")

	r = round{
		them: paper,
		goal: draw,
	}
	assert.Equal(t, r.score(), int(draw)+int(paper), "draw+paper != 5")

	r = round{
		them: scissors,
		goal: draw,
	}
	assert.Equal(t, r.score(), int(draw)+int(scissors), "draw+scissors != 6")

	r = round{
		them: scissors,
		goal: win,
	}
	assert.Equal(t, r.score(), int(win)+int(rock), "win+rock != 7")

	r = round{
		them: rock,
		goal: win,
	}
	assert.Equal(t, r.score(), int(win)+int(paper), "win+paper != 8")

	r = round{
		them: paper,
		goal: win,
	}
	assert.Equal(t, r.score(), int(win)+int(scissors), "win+scissors != 9")
}

func TestRound_update(t *testing.T) {
	want := round{
		them: paper,
		goal: win,
	}

	s := []move{paper, move(win)}
	got := new(round)
	got.update(s)

	assert.Equal(t, want.them, got.them)
	assert.Equal(t, want.goal, got.goal)
}
