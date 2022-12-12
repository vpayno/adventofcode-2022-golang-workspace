// Package day02 is the module with the cli logic for the cmd application.
package day02

// Config holds the application's configuration.
type Config struct {
	appName       string
	inputFileName string
}

// record ids (columns in input file)
const (
	them int = iota
	goal
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
}

var string2outcome = map[string]outcome{
	"X": loose,
	"Y": draw,
	"Z": win,
}

// Round describes one round of rock-paper-scissors.
type round struct {
	them move
	goal outcome
}

func (r *round) yourMove() move {
	var you move

	switch r.goal {
	case loose:
		switch r.them {
		case rock:
			you = scissors
		case paper:
			you = rock
		case scissors:
			you = paper
		}
	case win:
		switch r.them {
		case rock:
			you = paper
		case paper:
			you = scissors
		case scissors:
			you = rock
		}
	default:
		you = r.them
	}

	return you
}

func (r *round) score() int {
	return int(r.goal) + int(r.yourMove())
}

func (r *round) update(slice []move) {
	r.them = slice[them]
	r.goal = outcome(int(slice[goal]))
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
