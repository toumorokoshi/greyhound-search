package greyhound

import "io/ioutil"
import "github.com/toumorokoshi/go-fuzzy/fuzzy"

type SearchIndex struct {
	Files []string
	Matcher fuzzy.Matcher 
}

func NewSearchIndex (rootDir string) *SearchIndex {
	entries, _ := ioutil.ReadDir(rootDir)
	l := len(entries)
	files := make([]string, l, l)
	for pos, entry := range entries {
		files[pos] = entry.Name()
	}
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
