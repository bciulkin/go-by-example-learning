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
	
	fmt.Println("Empty input:", []Item{})

	result2, resultValue2 := knapstack([]Item{}, value_provider.IntN(70))
	fmt.Println("Result:", result2, resultValue2)
}

type Item struct {
	weight int
	value int
}


func knapstack(items []Item, weightLimit int) (int, int) {
	resultValue := 0
	resultWeight := 0

	n := len(items)
	dp := make([][]int, n+1)

	
	// if limit is 0 result is zero
	if weightLimit == 0 {
		return resultWeight, resultValue
	}

	if len(items) == 0 {
		return resultWeight, resultValue
	}

	// sum all items to have max value regardless weight limit
	for i := 0; i < len(items); i++ {
		resultWeight += items[i].weight
		resultValue += items[i].value
	}

	// as long as result weight is bigger weight limit, reduce one item
	for weightLimit > resultWeight {
		fmt.Println("current resultWeight: ", resultWeight)
		maxValue := resultValue - items[0].value
		for j := 1; j < len(items); j++ {
			maxValue := maxOf(maxValue,	resultValue - item[j].value)
		}
	}



	return resultWeight, resultValue
}

func maxOf(a int, b int) int {
	if a > b {
		return a
	else {
		return b
	}
}
