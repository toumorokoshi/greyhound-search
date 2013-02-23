package greyhound

import "io/ioutil"
import "encoding/json"

type SearchIndex struct {
	Files []string
}

func NewSearchIndex (rootDir string) *SearchIndex {
	entries, _ := ioutil.ReadDir(rootDir)
	l := len(entries)
	files := make([]string, l, l)
	for i := 0; i < l; i++ {
		files[i] = entries[i].Name()
	}
	return &SearchIndex{files}
}

// return the results for a search string m with a json result string
func (si *SearchIndex) ResultsJson(query string) string {
	result, _ := json.Marshal(si.Files)
	return string(result)
}
