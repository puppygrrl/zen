package game

import (
	"github.com/puppygrrl/zen/dictionary"
)

const maxGuesses int = 6

type Game struct {
	dict *dictionary.Dictionary
	guesses int
	word string
}

func New() *Game {
	newGame := Game{
		dict: dictionary.New(),
		guesses: maxGuesses,
	}
	newGame.word = newGame.dict.GetRandomCommonWord()
	return &newGame
}

func (g *Game) Word() string {
	return g.word
}
