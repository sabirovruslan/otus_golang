package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(input string) []string {
	result := make([]string, 0)
	if len(input) < 1 {
		return result
	}

	wordMap := make(map[string]int)
	for _, word := range strings.Fields(input) {
		if len(word) == 0 {
			continue
		}
		wordMap[word]++
	}

	wordByFrequencyMap := make(map[int][]string)
	for word, frequency := range wordMap {
		wordByFrequencyMap[frequency] = append(wordByFrequencyMap[frequency], word)
	}

	keys := make([]int, 0)
	for frequency := range wordByFrequencyMap {
		keys = append(keys, frequency)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, frequency := range keys {
		words := wordByFrequencyMap[frequency]
		sort.Strings(words)
		result = append(result, words...)
	}

	if len(result) > 10 {
		result = result[:10]
	}

	return result
}
