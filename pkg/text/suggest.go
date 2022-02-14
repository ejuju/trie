package text

import (
	"sort"
)

// Suggest returns a list of strings that could be the end of the given input
func (t *Trie) Suggest(str string, limit int) ([]string, error) {
	startNode, err := getLastNode(t, str)
	if err != nil {
		return nil, err
	}

	var out []string
	t.getSuggestions(startNode, "", &out)

	// post processing
	out = limitResults(limit, out)
	sortByLength(out)

	return out, nil
}

func (t *Trie) getSuggestions(root *TrieNode, prefix string, results *[]string) {
	if root == nil {
		return
	}

	for i, node := range root.Children {
		if node == nil {
			continue
		}

		newStr := prefix + string(t.Chars[i])

		if node.IsEnd {
			*results = append(*results, newStr)
		}

		t.getSuggestions(node, newStr, results)
	}
}

func limitResults(limit int, results []string) []string {
	if len(results) > limit {
		return results[:limit]
	}
	return results
}

type stringSorter struct {
	strings []string
}

func (ss stringSorter) Len() int {
	return len(ss.strings)
}

func (ss stringSorter) Less(i, j int) bool {
	return len(ss.strings[i]) < len(ss.strings[j])
}

func (ss stringSorter) Swap(i, j int) {
	ss.strings[i], ss.strings[j] = ss.strings[j], ss.strings[i]
}

func sortByLength(in []string) []string {
	sort.Sort(stringSorter{
		strings: in,
	})
	return in
}
