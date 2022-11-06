package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

func Top10(input string) []string {
	result := make([]string, 0)
	if len(input) < 1 {
		return result
	}

	frequencyWordMap := make(map[string]int, 0)
	for _, word := range strings.Fields(input) {
		if len(word) == 0 {
			continue
		}
		frequencyWordMap[word]++
	}

	for key, i := range frequencyWordMap {
		fmt.Println(key, " = ", i)
		result = append(result, key)
	}
	fmt.Println("-------------")
	sort.Strings(result)
	sort.Slice(result, func(i, j int) bool {
		return frequencyWordMap[result[i]] > frequencyWordMap[result[j]]
	})

	if len(result) > 10 {
		result = result[:10]
	}

	return result
}
