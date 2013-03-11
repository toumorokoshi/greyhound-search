package greyhound

import "io/ioutil"
import "strings"

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

// return a string slice for the results for a search string m with a json result string
func (si *SearchIndex) Results(query string) []string {
  returnFiles := make([]string, 0, len(si.Files))
  for _, f := range si.Files {
      if(strings.Contains(f, query)) {
          returnFiles = append(returnFiles, f)
      }
  }
  return returnFiles
}
