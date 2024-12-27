package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Ngram struct {
	key string
	count int
	freq float64
}

func main() {
	input := os.Args
	args := input[1:]
	var err error
	n := 5
	limit := 50
	switch len(args) {
	case 1:
		break
	case 2:
		limit, err = strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("Limit number wasn't well-formed")
		}

	default:
		fmt.Println("Usage: freq [limit] <file>")
		log.Fatal()
	}

	filepath := args[0]
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal()
	}

	// lists := [][]Ngram{ {}, {}, {}, {}, {} }
	dics := []map[string]int { {}, {}, {}, {}, {} }
	lists := make([][]Ngram, n)
	lens := make([]int, n)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			word = strings.ToLower(word)

			for i := 0; i < n; i++ {
				getNgrams(dics[i], word, i+1)
				lens[i] += len(word) -i
			}
		}
	}

	for i := range dics {
		for k, v := range dics[i] {
			f := float64(v) / float64(lens[i])*100.0
			ngram := Ngram{ k, v, f, }
			lists[i] = append(lists[i], ngram)
		}
	}

	for _, list := range lists {
		sort.Slice(list, func(i, j int) bool {
			return list[i].count > list[j].count
		})
	}

	for rank := range lists[3] {
		if rank >= limit {
			break
		}
		fmt.Printf("\n%2d: ", rank+1)
		for i := range lists {
			item := lists[i][rank]
			fmt.Printf("%-5s %-6.2f | ", item.key, item.freq)
		}
	}
}

func getNgrams(m map[string]int, word string, n int) {
	runes := []rune(word)
	for i := 0; i < len(runes)-n+1; i++ {
		key := string(runes[i:i+n])
		m[key]++
	}
}


