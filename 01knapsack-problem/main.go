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
	limit1 := value_provider.IntN(50)

	fmt.Println("01Knapstack problem.")
	fmt.Println("Input: ", input, "limit: ", limit1)

	result := knapsack(input, limit1)
	fmt.Println("Result:", result)
	
	fmt.Println("Empty input:", []Item{})

	result2 := knapsack([]Item{}, value_provider.IntN(70))
	fmt.Println("Result:", result2)
	items := []Item{
        {weight: 2, value: 3},
        {weight: 3, value: 4},
        {weight: 4, value: 5},
        {weight: 5, value: 6},
	}
	capacity := 5
	fmt.Println("Maximum value:", knapsack(items, capacity)) // Output: 7
}

type Item struct {
	weight int
	value int
}


func knapsack(items []Item, weightLimit int) int {
	n := len(items)
	dp := make([]int, weightLimit+1)

	for i := 0; i < n; i++ {
		for w:= weightLimit; w >= items[i].weight; w-- {
			dp[w] = maxOf(dp[w], dp[w-items[i].weight]+items[i].value)
		}
	}

	fmt.Println(dp)
	return dp[weightLimit]
}

func maxOf(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
