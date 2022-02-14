package text

// Trie is a data structure that stores characters (helpful for autocompletion)
type Trie struct {
	Root  *TrieNode
	Chars string
}

// TrieNode represents a character in the trie data structure
type TrieNode struct {
	Children [NumPossibleChars]*TrieNode
	IsEnd    bool // IsEnd == true means that this is the last letter of a word
}

// NewTrie is a helper function to easily create a new trie
func NewTrie(strs []string) (*Trie, error) {
	t := &Trie{
		Root: &TrieNode{},
	}

	err := t.Fill(strs)
	return t, err
}

// AddString inserts a new string in the trie
func (t *Trie) addString(str string) error {
	currNode := t.Root
	for _, c := range str {
		index, ok := t.charToIndex(c)
		if ok == false {
			continue
		}
		if currNode.Children[index] == nil {
			currNode.Children[index] = &TrieNode{} // add node for this char if not defined yet
		}
		currNode = currNode.Children[index]
	}

	currNode.IsEnd = true
	return nil
}

// hasChildren iterates over all children of a node and returns false if no children is defined
func (n *TrieNode) hasChildren() bool {
	out := false
	for _, childNode := range n.Children {
		if childNode != nil {
			out = true
		}
	}
	return out
}
