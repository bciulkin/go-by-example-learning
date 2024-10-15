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

func lcs(s1, s2 string) int {
	m := len(s1)
	n := len(s2)

	// Create a 2D slice to store lengths of LCS for substrings
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	
    	for i := 1; i <= m; i++ {
        	for j := 1; j <= n; j++ {
			//fmt.Println("s1: ", s1[i-1])
			//fmt.Println("s2: ", s2[j-1])
			if s1[i-1] == s2[j-1] {
                		dp[i][j] = dp[i-1][j-1] + 1
			} else {
                		dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp)
	return dp[m][n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
