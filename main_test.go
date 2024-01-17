package main

import "testing"

// Benchmark for checkWords function
func BenchmarkCheckWords(b *testing.B) {
	words := loadDictionary("words.txt")
	letters := "ovmefshegxvytpnrqkxgyceosjyvcfhhhkynpdomgfwsmcmcqremqyosqhtvbuwkyypycjcfvrzupwpsqwvqvbxgftytxdkkpokm"
	results := make(chan string, len(words))

	for i := 0; i < b.N; i++ {
		results = make(chan string, len(words)) // Reset channel for each iteration
		go checkWords(letters, words, results)

		possibleWords := []string{}
		for word := range results {
			possibleWords = append(possibleWords, word)
		}
	}
}
