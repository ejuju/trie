package trie

import (
	"fmt"
	"testing"
	"time"
)

//
func TestLoad(t *testing.T) {
	strs, err := Load("./words/", "test.json")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if len(strs) != 1 {
		t.Log("result from trie.Load(\"./words/\", \"test.json\") should be of length 1")
		t.Fail()
		return
	}
}

func TestNewTrie(t *testing.T) {
	uniqueChars := "abcdefghiklmnopqrstuvwxyz0123456789- " // all characters in this string should be unique

	trie, err := New(uniqueChars)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if trie.Chars != uniqueChars {
		t.Log("trie.Chars should be equal to unique chars, expected: \"" + uniqueChars + "\" but got \"" + trie.Chars + "\"")
		t.Fail()
		return
	}
}

func TestSuggest(t *testing.T) {
	trie, err := New("a", "b", "abba")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	res, err := trie.Suggest("", 1)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	if len(res) != 1 {
		t.Log("trie.Suggest() should have returned 1 result")
		t.Fail()
		return
	}

	// todo: finish
}

func TestLong(t *testing.T) {
	startMain := time.Now()
	fmt.Println(">> Launching...")

	strs, err := Load("./words/", FrenchWords, EnglishWords)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	fmt.Printf(">> Loaded %v strings in %v milliseconds \n", len(strs), time.Now().Sub(startMain).Milliseconds())
	startFill := time.Now()

	trie, err := New(strs...)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	fmt.Printf(">> Initiated trie in %v milliseconds \n", time.Now().Sub(startFill).Milliseconds())

	startComplete := time.Now()

	input := "a"
	limit := 50
	res, err := trie.Suggest(input, limit)

	if err != nil {
		t.Log(">> Error:", err)
		t.Fail()
		return
	}

	fmt.Printf(">> Found %v possible results in %v milliseconds\n", len(res), time.Now().Sub(startComplete).Milliseconds())

	if res != nil {
		fmt.Printf(">> %v results for \"%v\" \n%v \n", len(res), input, res)
	}
}
