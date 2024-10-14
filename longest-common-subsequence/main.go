package main

import (
	"fmt"
	"github.com/bciulkin/go-test-value-provider"
)

func main() {
	input := value_provider.String()
	input2 := value_provider.String()

	fmt.Println("Longest common subsequence.")
	fmt.Println("Input: ", input, "input2: ", input2)

	result := lcs(input, input2)
	fmt.Println("Result: ", result)


	input3 := "ABCDEFG"
	input4 := "DEGJABCDAE"

	fmt.Println("Another example.")
	fmt.Println("Input: ", input3, "input2: ", input4)

	result2 := lcs(input3, input4)
	fmt.Println("Result: ", result2) // should be ABCD
	
}

func lcs(s1, s2 string) string {
	return "ABC"
}


