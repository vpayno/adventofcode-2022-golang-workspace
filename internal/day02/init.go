// Package day02 is the module with the cli logic for the cmd application.
package day02

// Config holds the application's configuration.
type Config struct {
	appName       string
	inputFileName string
}

const (
	them int = iota
	you
)

// Outcome describes the possible outcomes.
type outcome int

const (
	loose outcome = 0
	draw  outcome = 3
	win   outcome = 6
)

var outcomeNames = map[outcome]string{
	loose: "loose",
	draw:  "draw",
	win:   "win",
}

// Move descrives possible moves.
type move int

const (
	rock     move = 1
	paper    move = 2
	scissors move = 3
)

var moveNames = map[move]string{
	rock:     "rock",
	paper:    "paper",
	scissors: "scissors",
}

var string2move = map[string]move{
	"A": rock,
	"B": paper,
	"C": scissors,
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

// Round describes one round of rock-paper-scissors.
type round struct {
	them move
	you  move
}

func (r *round) judge() outcome {
	if r.you == r.them {
		return draw
	}

	if r.you == rock && r.them == scissors {
		return win
	}

	if r.you == paper && r.them == rock {
		return win
	}

	if r.you == scissors && r.them == paper {
		return win
	}

	return loose
}

func (r *round) score() int {
	return int(r.judge()) + int(r.you)
}

func (r *round) update(slice []move) {
	r.them = slice[them]
	r.you = slice[you]
}

type rounds []round

// Setup creates the applications configuration object.
func Setup(appName string) Config {

	conf := Config{
		appName:       appName,
		inputFileName: "data/" + appName + "/" + appName + "-input.txt",
	}

	return conf
}
