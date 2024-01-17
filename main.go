package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func loadDictionary(filename string) []string {
	words := []string{}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		words = append(words, word)
	}
	return words
}

func checkWords(letters string, words []string, results chan string) {
	letterCounter := make(map[rune]int)
	for _, r := range strings.ToLower(letters) {
		letterCounter[r]++
	}

	var wg sync.WaitGroup

	for _, word := range words {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()

			wordCounter := make(map[rune]int)
			for _, r := range word {
				wordCounter[r]++
			}

			canConstruct := true
			for r, count := range wordCounter {
				if letterCounter[r] < count {
					canConstruct = false
					break
				}
			}

			if canConstruct {
				results <- word
			}
		}(word)
	}

	go func() {
		wg.Wait()
		close(results)
	}()
}

func main() {
	words := loadDictionary("words.txt")
	letters := "shubhamkoli"

	results := make(chan string, len(words)) // Buffered channel

	go checkWords(letters, words, results)

	possibleWords := []string{}
	for word := range results {
		possibleWords = append(possibleWords, word)
	}

	fmt.Println("Total ", len(possibleWords), "Words made from string! Check them out \n ", possibleWords)
}
