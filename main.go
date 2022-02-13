package main

import (
	"fmt"
	"time"

	"github.com/ejuju/trie-implementation-autocomplete/pkg/trie"
)

func main() {
	startMain := time.Now()
	fmt.Println(">> Launching...")

	t := trie.New()

	dict, err := trie.DefaultDictionary()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(">> Loaded default dictionary (%v words) in %v milliseconds \n", len(dict), time.Now().Sub(startMain).Milliseconds())
	startFill := time.Now()

	err = t.Fill(dict)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(">> Filled trie with default dictionary (%v words) in %v milliseconds \n", len(dict), time.Now().Sub(startFill).Milliseconds())
	startComplete := time.Now()

	input := "r"
	limit := 50
	res, err := t.Complete(input, limit)

	if err != nil {
		fmt.Println(">> Error:", err)
		return
	}

	fmt.Printf(">> Found %v possible results in %v milliseconds\n", len(res), time.Now().Sub(startComplete).Milliseconds())

	if res != nil {
		fmt.Printf(">> %v results for \"%v\" \n%v \n", len(res), input, res)
	}
}
