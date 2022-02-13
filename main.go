package main

import (
	"fmt"
	"time"

	"github.com/ejuju/trie-implementation-autocomplete/pkg/trie"
)

func main() {
	start := time.Now()
	fmt.Println("Launching...")

	t := trie.New()

	dict, err := trie.DefaultDictionary()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = t.Fill(dict)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Filled trie with default dictionary in", time.Now().Sub(start).Seconds(), "seconds")
	fmt.Println("Default dictionary has", len(dict), "words")

	input := "z"
	maxResults := 10

	res, err := t.Complete(input, maxResults)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if res != nil {
		fmt.Println(len(res), "results:", res)
		return
	}

	fmt.Println("Done in", time.Now().Sub(start).Seconds(), "seconds")
}
