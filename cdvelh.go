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
		if input <= diff {
			return 0
		}
		if input%diff != 0 {
			return 0
		}
		result = input / diff
	case OperationMod:
		if input <= diff {
			return 0
		}
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

func Shuffle(src []string) []string {
	dest := make([]string, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}

func (b *Book) PrintMarkdown() {
	fmt.Println("% La calculatrice dont vous êtes le héros.")
	fmt.Println("% Manfred Touron")

	for i := 0; i < b.Pages; i++ {
		chiffre := b.Mapping[i] + 1
		//page := i + b.Base

		totalOperations := []string{}

		for operation := OperationAdd; operation <= OperationMod; operation++ {
			operationLines := []string{}
			for diff := 1; diff < b.Pages; diff++ {
				result := b.Operation(chiffre, operation, diff)
				if result != 0 && result != chiffre {
					symbol := SymbolMapping[operation]
					operationLines = append(operationLines, fmt.Sprintf("| **%d %s %d** | voir page %d |", chiffre, symbol, diff, b.GetPage(result)))
				}
			}

			for i, line := range Shuffle(operationLines) {
				if i > 5 {
					break
				}
				totalOperations = append(totalOperations, line)
			}
		}

		// PRINT
		fmt.Printf("# %d\n", chiffre)
		fmt.Println("")
		fmt.Println("| Calcul | Solution |")
		fmt.Println("| :-- | --: |")
		// fmt.Printf("Liste des calculs de pro avec %d:\n", chiffre)
		//fmt.Println("")

		i := 0
		for _, line := range Shuffle(totalOperations) {
			if i > 10 {
				break
			}
			fmt.Println(line)
			i++
		}

		fmt.Println("")
		//fmt.Printf("Page %d\n", page)
		fmt.Println("")
		fmt.Println("---")
		fmt.Println("")
	}
}
