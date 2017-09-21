package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}
	// scan the response body
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// create slice to hold counts
	buckets := make([]int, 200)
	// loop over words
	for scanner.Scan() {
		n := hashBucket(scanner.Text())
		buckets[n]++
	}
	fmt.Println(buckets[65:122])
}

func hashBucket(word string) int {
	return int(word[0])
}
