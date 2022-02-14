package text

import (
	"encoding/json"
	"os"
)

// Load returns a list of words to fill the trie with
func Load(filepaths ...string) ([]string, error) {
	var out []string

	pathPrefix := "./pkg/text/"

	for _, fp := range filepaths {
		f, err := os.Open(pathPrefix + fp)
		if err != nil {
			return nil, err
		}

		var words []string
		err = json.NewDecoder(f).Decode(&words)
		if err != nil {
			return nil, err
		}

		out = append(out, words...)
	}

	return out, nil
}