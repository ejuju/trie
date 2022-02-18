package main

import (
	"fmt"
	"time"

	"github.com/ejuju/trie-implementation-autocomplete/pkg/trie"
)

func main() {
	startMain := time.Now()
	fmt.Println(">> Launching...")

	strs, err := trie.Load(trie.FrenchWords, trie.EnglishWords)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(">> Loaded %v strings in %v milliseconds \n", len(strs), time.Now().Sub(startMain).Milliseconds())
	startFill := time.Now()

	t, err := trie.New(strs...)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(">> Initiated trie in %v milliseconds \n", time.Now().Sub(startFill).Milliseconds())

	startComplete := time.Now()

	input := "a"
	limit := 50
	res, err := t.Suggest(input, limit)

	if err != nil {
		fmt.Println(">> Error:", err)
		return
	}

	fmt.Printf(">> Found %v possible results in %v milliseconds\n", len(res), time.Now().Sub(startComplete).Milliseconds())

	if res != nil {
		fmt.Printf(">> %v results for \"%v\" \n%v \n", len(res), input, res)
	}
}
