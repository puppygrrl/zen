package main

import (
	"fmt"
	"github.com/puppygrrl/zen/game"
)

func main() {
	g := game.New()
	fmt.Println(g.Word())
}
