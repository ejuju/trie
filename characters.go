package trie

import (
	"strings"
)

// MaxChildrenPerNode limits the number of children a trie node can have.
// It corresponds to the maximum number of characters allowed (1 child = 1 character)
const MaxChildrenPerNode = 256

// Words for various languages
const (
	englishWords = "words_en.json" // src: https://github.com/words/an-array-of-english-words
	frenchWords  = "words_fr.json" // src: https://github.com/words/an-array-of-french-words
)

// fill is utility function that wraps t.Add() to add a list of strings to the trie
func (t *Trie) fill(in []string) error {
	var err error

	for _, s := range in {
		for _, c := range s {
			if !strings.ContainsRune(t.Chars, c) {
				t.Chars += string(c)
			}
		}

		err = t.addString(s)
		if err != nil {
			return err
		}
	}

	return nil
}

// charToIndex converts a rune to the given index (of the trie node children).
// It returns false if the character is not allowed
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
