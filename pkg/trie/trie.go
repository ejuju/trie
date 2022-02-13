package trie

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Alphabet represents possible characters (a to z, whitespace, 0 to 9, special characters like the dash "-") used in tries
const Alphabet = "abcdefghijklmnopqrstuvwxyz0123456789 -"

// AlphabetSize holds the number of possible characters in the alphabet
const AlphabetSize = len(Alphabet)

// Trie is a data structure that stores characters (helpful for autocompletion)
type Trie struct {
	Root *Node
}

// Node represents a character in the trie data structure
type Node struct {
	Children [AlphabetSize]*Node
	IsEnd    bool // IsEnd == true means that this is the last letter of a word
}

// CharToIndex holds a map of character to an index of the trie node's children
var CharToIndex = initCharToIndex(Alphabet)

// IndexToChar converts an index to the corresponding character of the alphabet
var IndexToChar = initIndexToChar(Alphabet)

//
func initCharToIndex(alphabet string) map[rune]int {
	result := map[rune]int{}
	for i, c := range alphabet {
		result[c] = i
	}
	return result
}

//
func initIndexToChar(alphabet string) map[int]rune {
	result := map[int]rune{}
	for i, c := range alphabet {
		result[i] = c
	}
	return result
}

// HasChildren iterates over all children of a node and returns false if no children is defined
func (n *Node) HasChildren() bool {
	out := false
	for _, childNode := range n.Children {
		if childNode != nil {
			out = true
		}
	}
	return out
}

// New is a helper function to easily create a new trie
func New() *Trie {
	return &Trie{Root: &Node{}}
}

// Add inserts a new word in the trie
func (t *Trie) Add(word string) error {
	currNode := t.Root
	for _, c := range word {
		index, ok := CharToIndex[c]
		if ok == false {
			return fmt.Errorf("%v is not allowed", string(c))
		}
		if currNode.Children[index] == nil {
			currNode.Children[index] = &Node{} // add node for this char if not defined yet
		}
		currNode = currNode.Children[index]
	}
	currNode.IsEnd = true
	return nil
}

// HasWord checks if a word is included in the trie
func (t *Trie) HasWord(word string) bool {
	currNode := t.Root
	for _, c := range word {
		index := CharToIndex[c]
		if currNode.Children[index] == nil {
			return false
		}
		currNode = currNode.Children[index]
	}
	return currNode.IsEnd
}

// Complete returns a list of strings that could be the end of the given input
func (t *Trie) Complete(str string, limit int) ([]string, error) {
	startNode, err := getLastNode(t, str)
	if err != nil {
		return nil, err
	}

	var results []string
	var resCh = make(chan string, limit)

	go searchRec(startNode, str, limit, resCh)

	for i := 0; i < limit; i++ {
		res := <-resCh
		results = append(results, res)
	}

	return results, nil
}

func searchRec(root *Node, prefix string, limit int, results chan string) {
	if root == nil {
		return
	}

	for i, node := range root.Children {
		newStr := prefix + string(IndexToChar[i])

		if node == nil {
			continue
		}

		if node.IsEnd {
			results <- newStr
		}

		go searchRec(node, newStr, limit, results)
	}
}

// getLastNode returns the last node (= the last character) for a given string
func getLastNode(t *Trie, str string) (*Node, error) {
	currNode := t.Root

	for i := 0; i < len(str); i++ {
		currChar := rune(str[i])
		charIndex := CharToIndex[currChar]

		if currNode.Children[charIndex] == nil {
			return nil, errors.New("unable to find last node of string, the provided string is not defined in the trie")
		}

		currNode = currNode.Children[charIndex] // update current node to match current char
	}

	return currNode, nil
}

// Fill is utility function that wraps t.Add() to add a list of strings to the trie
func (t *Trie) Fill(in []string) error {
	var err error
	for _, s := range in {
		err = t.Add(s)
		if err != nil {
			return err
		}
	}
	return nil
}

// DefaultDictionary returns a list of strings to fill the trie with
// Words come from the following list: https://github.com/dwyl/english-words/blob/master/words_dictionary.json which contains around 300K words
func DefaultDictionary() ([]string, error) {
	f, err := os.Open("./pkg/trie/words_dictionary.json")
	if err != nil {
		return nil, err
	}

	var dictionary map[string]int

	err = json.NewDecoder(f).Decode(&dictionary)
	if err != nil {
		return nil, err
	}

	var out []string
	for key := range dictionary {
		out = append(out, key)
	}

	return out, nil
}
