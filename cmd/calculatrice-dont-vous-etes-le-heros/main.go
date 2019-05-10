package main

import (
	"math/rand"
	"time"

	"github.com/ultreme/calculatrice-dont-vous-etes-le-heros"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	book := cdvelh.MakeBook(1000, 3)
	book.PrintMarkdown()
}
