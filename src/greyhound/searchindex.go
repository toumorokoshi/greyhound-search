package greyhound

import "io/ioutil"
import "io"
import "github.com/toumorokoshi/go-fuzzy/fuzzy"
import "fmt"

type SearchIndex struct {
	Files []string
	Matcher fuzzy.Matcher 
}

func recursiveSearch (filePaths *[]string, file *io.File, prefix string) {
	if(!file.Mode().isDir()) {
		filePaths = &append(*filePaths, file.Name())
	} else {
		prefix = fmt.SprintF("%s/%s", prefix, file.Name())
		entries, _ := ioutil.ReadDir(file.Name())
		for _, entry := range entries {
			recursiveSearch(filePaths, entry, prefix)
		}
	}
}

func NewSearchIndex (rootDir string) *SearchIndex {
	files := make([]string, 0, 10000)
	recursiveSearch(&file, io.Open(rootDir), "")
	return &SearchIndex{files, fuzzy.NewMatcher(files)}
}
// return a string slice for the results for a search string m with a json result string
func (si *SearchIndex) Results(query string) []string {
	matches := si.Matcher.ClosestList(query, 5)
	matchStrings := make([]string, len(matches))
	for pos, value := range matches {
		matchStrings[pos] = value.Value
	}
	return matchStrings
}
