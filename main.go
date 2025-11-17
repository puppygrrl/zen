package main

import (
	"github.com/puppygrrl/zen/dictionary"
)

func main() {
	d := dictionary.New()
	fmt.Println(d.GetRandomCommonWord())
}
