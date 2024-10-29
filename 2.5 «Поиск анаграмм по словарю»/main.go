package main

import (
	"fmt"
	"sort"
	"strings"
)

// sortWord сортирует буквы в слове в алфавитном порядке.

func sortWord(word string) string {
	runes := []rune(word)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// findAnagram принимает массив слов, приводит каждое слово к нижнему регистру,
// сортирует буквы в слове, и использует отсортированное представление как ключ
// для группировки слов по анаграммам. Результат — мапа, где ключом является
// первое слово из группы анаграмм, а значением — отсортированный массив анаграмм.

func findAnagram(words []string) map[string][]string {
	result := make(map[string][]string)
	anagrams := make(map[string][]string)

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		sortedWord := sortWord(lowerWord)
		anagrams[sortedWord] = append(anagrams[sortedWord], lowerWord)
	}

	for _, group := range anagrams {
		sort.Strings(group)
		result[group[0]] = group
	}
	return result
}

func main() {
	words := []string{"слиток", "пятак", "пятка", "тяпка", "листок", "столик"}
	result := findAnagram(words)
	fmt.Println(result)
}
