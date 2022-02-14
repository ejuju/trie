package text

import (
	"strings"
)

//
const MaxChildrenPerNode = 256

// Words for various languages
const (
	EnglishWords = "words_en.json" // src: https://github.com/words/an-array-of-english-words
	FrenchWords  = "words_fr.json" // src: https://github.com/words/an-array-of-french-words
)

// Fill is utility function that wraps t.Add() to add a list of strings to the trie
func (t *Trie) fill(in []string) error {
	var err error

	for _, str := range in {
		for _, c := range str {
			if strings.ContainsRune(t.Chars, c) {
				continue
			}
			t.Chars += string(c)
		}
	}

	for _, s := range in {
		err = t.addString(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Trie) charToIndex(c rune) (int, bool) {
	i := strings.IndexRune(t.Chars, c)
	ok := i != -1
	return i, ok
}

// HasWord checks if a word is included in the trie
func (t *Trie) HasWord(word string) bool {
	currNode := t.Root
	for _, c := range word {
		index, ok := t.charToIndex(c)
		if ok == false {
			return false
		}

		if currNode.Children[index] == nil {
			return false
		}
		currNode = currNode.Children[index]
	}
	return currNode.IsEnd
}
