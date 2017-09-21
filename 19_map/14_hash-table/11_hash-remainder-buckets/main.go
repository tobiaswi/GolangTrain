package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	buckets := make([][]string, 12)
	// create slices to hold words
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}

	// loop over the words
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, len(buckets))
		buckets[n] = append(buckets[n], word)
	}

	// print the length of each bucket
	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}

	// print the words in one of the buckets
	// fmt.Println(buckets[6])
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
