package cdvelh

import (
	"fmt"
	"math/rand"
)

type Book struct {
	Base    int
	Mapping []int
	Pages   int
}

type Page struct{}

var SymbolMapping map[int]string

func init() {
	SymbolMapping = map[int]string{
		OperationAdd: "+",
		OperationSub: "-",
		OperationMul: "*",
		OperationDiv: "/",
		OperationMod: "modulo",
	}
}

func MakeBook(pages int, base int) Book {
	book := Book{
		Base:    base,
		Mapping: rand.Perm(pages),
		Pages:   pages,
	}
	return book
}

const (
	OperationAdd = iota
	OperationSub
	OperationMul
	OperationDiv
	OperationMod
)

func (b *Book) Operation(input int, operation int, diff int) int {
	var result int
	switch operation {
	case OperationAdd:
		result = input + diff
	case OperationSub:
		result = input - diff
	case OperationMul:
		result = input * diff
	case OperationDiv:
		result = input / diff
	case OperationMod:
		result = input % diff
	}

	if result <= 1 {
		return 0
	}
	if result >= b.Pages {
		return 0
	}
	return result
}

func (b *Book) GetPage(number int) int {
	for i := 0; i < b.Pages; i++ {
		if b.Mapping[i] == number-1 {
			return i + b.Base
		}
	}
	return 0
}

func (b *Book) PrintMarkdown() {
	fmt.Println("# La calculatrice dont vous êtes le héros.")

	for i := 0; i < b.Pages; i++ {
		chiffre := b.Mapping[i] + 1
		//page := i + b.Base
		fmt.Printf("# Le chiffre %d\n", chiffre)
		fmt.Println("")

		for operation := OperationAdd; operation <= OperationMod; operation++ {
			for diff := 1; diff < b.Pages; diff++ {
				result := b.Operation(chiffre, operation, diff)
				if result != 0 && result != chiffre {
					symbol := SymbolMapping[operation]
					fmt.Printf("* Si tu veux voir combien font %d %s %d, rends toi page %d\n", chiffre, symbol, diff, b.GetPage(result))
				}
			}
		}

		fmt.Println("")
		//fmt.Printf("Page %d\n", page)
		fmt.Println("")
		fmt.Println("---")
		fmt.Println("")
	}
}
