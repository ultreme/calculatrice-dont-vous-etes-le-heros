package main

import (
	"math/rand"
	"time"

	"github.com/ultreme/calculatrice-dont-vous-etes-le-hero"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	book := cdvelh.MakeBook(1000, 2)
	book.PrintMarkdown()
}
