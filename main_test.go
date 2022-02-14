package main

import (
	"testing"

	"github.com/ejuju/trie-implementation-autocomplete/pkg/text"
)

//
func TestLoad(t *testing.T) {
	strs, err := text.Load("test.json")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if len(strs) != 1 {
		t.Log("result from text.Load(\"test.json\") should be of length 1")
		t.Fail()
		return
	}
}

func TestNewTrie(t *testing.T) {
	uniqueChars := "abcdefghiklmnopqrstuvwxyz0123456789- " // all characters in this string should be unique

	trie, err := text.NewTrie(uniqueChars)
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
	trie, err := text.NewTrie("a", "b", "abba")
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
