package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Ngram struct {
	key string
	count int
	freq float64
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal()
	}
	filepath := args[0]
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal()
	}
	charCount := 0
	wordCount := 0
	n2Count := 0
	n3Count := 0

	charMap := map[string]int{}
	biMap := map[string]int{}
	triMap := map[string]int{}
	wordMap := map[string]int{}

	// maps := []map[string]int {
	// 	charMap,
	// 	biMap,
	// 	triMap,
	// 	wordMap,
	// }

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			wordCount++
			wordMap[word]++
			for _, c := range word {
				charCount++
				charMap[strings.ToLower(string(c))]++
			}
			getNgrams(biMap, word, 2)
			getNgrams(triMap, word, 3)
			n2Count += len(word) - 1
			n3Count += len(word) - 2
		}
	}
	chars := []Ngram{}
	words := []Ngram{}
	n2grams := []Ngram{}
	n3grams := []Ngram{}

	for k, v := range charMap {
		chars = append(chars, Ngram{
			string(k),
			v,
			float64(v)/float64(charCount)*100,
		})
	}
	for k, v := range wordMap {
		words = append(words, Ngram{
			string(k),
			v,
			float64(v)/float64(wordCount)*100,
		})
	}
	for k, v := range biMap {
		n2grams = append(n2grams, Ngram{
			string(k),
			v,
			float64(v)/float64(n2Count)*100,
		})
	}
	for k, v := range triMap {
		n3grams = append(n3grams, Ngram{
			string(k),
			v,
			float64(v)/float64(n3Count)*100,
		})
	}
	sort.Slice(chars, func(i, j int) bool {
		return chars[i].count > chars[j].count
	})
	sort.Slice(words, func(i, j int) bool {
		return words[i].count > words[j].count
	})
	sort.Slice(n2grams, func(i, j int) bool {
		return n2grams[i].count > n2grams[j].count
	})
	sort.Slice(n3grams, func(i, j int) bool {
		return n3grams[i].count > n3grams[j].count
	})

	for i, ngram := range chars {
		fmt.Printf("%2d: ", i+1)
		fmt.Printf("%-5s %-6.2f | ", ngram.key, ngram.freq)
		fmt.Printf("%-5s %-6.2f | ", n2grams[i].key, n2grams[i].freq)
		fmt.Printf("%-5s %-6.2f | ", n3grams[i].key, n3grams[i].freq)
		fmt.Printf("%-16s %-6.2f\n", words[i].key, words[i].freq)
		if i > 20 {
			break
		}
	}
}

func getNgrams(m map[string]int, word string, n int) {
	runes := []rune(word)
	for i := 0; i < len(runes)-n+1; i++ {
		m[string(runes[i:i+2])]++
	}
}

