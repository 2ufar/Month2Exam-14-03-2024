package main

import (
	"fmt"
)

func longestCommonPrefix(strings []string) string {
	if len(strings) == 0 {
		return "There is no word to prefix"
	}

	shortest := strings[0]
	for _, s := range strings {
		if len(s) < len(shortest) {
			shortest = s
		}
	}
	for i := 0; i < len(shortest); i++ {
		for j := 1; j < len(strings); j++ {
			if strings[j][i] != shortest[i] {
				return shortest[:i]
			}
		}
	}
	return shortest
}

func main() {
	strings1 := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strings1))

	strings2 := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strings2)) 
}
