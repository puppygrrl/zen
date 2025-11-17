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

func Guess(word string, secret string) string {
	if word == secret {
		return ""
	}
	hint := ""
	for pos, char := range word {
		if secret[pos] == byte(char) {
			hint += "1" // green
		} else if strings.Contains(secret, string(char)) {
			hint += "2" // yellow
		} else {
			hint += "0" // grey
		}
	}

	return hint
}
