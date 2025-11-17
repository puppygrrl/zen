package game

import (
	"github.com/puppygrrl/zen/dictionary"
	"strings"
)

const maxGuesses int = 6

type Game struct {
	dict    *dictionary.Dictionary
	guesses int
	word    string
}

func New() *Game {
	newGame := Game{
		dict:    dictionary.New(),
		guesses: maxGuesses,
	}
	newGame.word = newGame.dict.GetRandomCommonWord()
	return &newGame
}

func (g *Game) Word() string {
	return g.word
}

func (g *Game) Guess(word string) string {
	if word == g.word {
		return ""
	}
	hint := ""
	for pos, char := range word {
		if g.word[pos] == byte(char) {
			hint += "1" // green
		} else if strings.Contains(word, string(char)) {
			hint += "2" // yellow
		} else {
			hint += "0" // grey
		}
	}

	return hint
}
