package trie

import "sort"

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

func limitResults(limit int, results []string) []string {
	return results[:limit]
}
