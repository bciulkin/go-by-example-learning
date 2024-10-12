package main

import (
	"fmt"
	"github.com/bciulkin/go-test-value-provider"
)

func main() {
	input := []Item {
		Item{weight: value_provider.IntN(5), value: value_provider.IntN(5)},
		Item{weight: value_provider.IntN(10), value: value_provider.IntN(10)},
		Item{weight: value_provider.IntN(20), value: value_provider.IntN(20)},
		Item{weight: value_provider.IntN(30), value: value_provider.IntN(30)},
		Item{weight: value_provider.IntN(40), value: value_provider.IntN(40)},
	}

	fmt.Println("01Knapstack problem.")
	fmt.Println("Input:", input)

	result, resultValue := knapstack(input, value_provider.IntN(50))
	fmt.Println("Result:", result, resultValue)
}

type Item struct {
	weight int
	value int
}


func knapstack(items []Item, weightLimit int) ([]Item, int) {
	var result []Item
	resultValue := 0

	return result, resultValue
}
