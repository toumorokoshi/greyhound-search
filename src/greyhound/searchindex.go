package greyhound

import (
	"github.com/toumorokoshi/go-fuzzy/fuzzy"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type SearchIndex struct {
	Files   []string
	Matcher fuzzy.Matcher
	Exclusions []*regexp.Regexp
}

func exclude(exclusions []*regexp.Regexp, target string) bool {
	for _, exclusion := range exclusions {
		if exclusion.MatchString(target) {
			return true
		}
	}
	return false
}

func recursiveSearch(filePaths []string, file os.FileInfo, prefix, root string, exclusions []*regexp.Regexp) []string {
	if !file.Mode().IsDir() {
		if !exclude(exclusions, file.Name()) {
			filePaths = append(filePaths, file.Name())
		}
	} else {
		dirPath := strings.Join([]string{root, prefix, file.Name()}, "/")
		prefix = strings.Join([]string{prefix, file.Name()}, "/")
		entries, err := ioutil.ReadDir(dirPath)
		if err != nil {
			log.Print("Error: unable to read directory: ", err.Error())
		}
		for _, entry := range entries {
			filePaths = recursiveSearch(filePaths, entry, prefix, root, exclusions)
		}
	}
	return filePaths
}

func NewSearchIndex(rootDir string, exclusions []*regexp.Regexp) *SearchIndex {
	files := make([]string, 0, 1000000)
	dir, err := os.Lstat(rootDir)
	if err != nil {
		log.Print("Error: unable to open root path: ", err.Error())
	}
	paths := strings.Split(rootDir, "/")
	rootDir = strings.Join(paths[0:len(paths) - 2], "/")
	files = recursiveSearch(files, dir, "", rootDir, exclusions)
	log.Print("Total Filecount: ", len(files))
	return &SearchIndex{files, fuzzy.NewMatcher(files), exclusions}
}

// return a string slice for the results for a search string m with a json result string
func (si *SearchIndex) Results(query string) []string {
	matches := si.Matcher.ClosestList(query, 20)
	matchStrings := make([]string, len(matches))
	for pos, value := range matches {
		matchStrings[pos] = value.Value
	}
	return matchStrings
}
