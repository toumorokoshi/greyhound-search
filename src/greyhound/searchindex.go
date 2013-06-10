package greyhound

import (
	"fmt"
	"github.com/toumorokoshi/go-fuzzy/fuzzy"
	"path/filepath"
	"code.google.com/p/codesearch/index"
	csregexp "code.google.com/p/codesearch/regexp"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type IndexedFile struct {
	FullPath, FileName string
}

type SearchIndex struct {
	Matcher fuzzy.Matcher
	CodeIndex index.Index
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

func recursiveSearch(filePaths []fuzzy.MatchStruct, 
	                   codesearchIndex *index.IndexWriter,
                   	 file os.FileInfo, 
                   	 prefix, root string, 
                     exclusions []*regexp.Regexp) []fuzzy.MatchStruct {
	path := strings.Join([]string{root, prefix, file.Name()}, "/")
	if !file.Mode().IsDir() {
		if !exclude(exclusions, file.Name()) {
			filePaths = append(filePaths, fuzzy.MatchStruct{file.Name(), map[string]string{"fullPath": path}})
			codesearchIndex.AddFile(path)
		}
	} else {
		prefix = strings.Join([]string{prefix, file.Name()}, "/")
		entries, err := ioutil.ReadDir(path)
		if err != nil {
			log.Print("Error: unable to read directory: ", err.Error())
		}
		for _, entry := range entries {
			filePaths = recursiveSearch(filePaths, codesearchIndex, entry, prefix, root, exclusions)
		}
	}
	return filePaths
}

func NewSearchIndex(name, rootDir string, exclusions []*regexp.Regexp) *SearchIndex {
	files := make([]fuzzy.MatchStruct, 0, 1000000)
	dir, err := os.Lstat(rootDir)
	if err != nil {
		log.Print("Error: unable to open root path: ", err.Error())
	}
	rootDir = strings.TrimSuffix(rootDir, "/")
	paths := strings.Split(rootDir, "/")
	rootDir = strings.Join(paths[0:len(paths) - 1], "/")
	codeindexPath, err := filepath.Abs(fmt.Sprintf("./indices/%s.ix", name))
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Create(codeindexPath); err != nil {
		log.Fatal(err)
	}
	codeindex := index.Create(codeindexPath)
	codeindex.AddPaths([]string{rootDir})
	files = recursiveSearch(files, codeindex, dir, "", rootDir, exclusions)
	codeindex.Flush()
	log.Print("Total Filecount: ", len(files))
	return &SearchIndex{fuzzy.NewMatcher(files), *index.Open(codeindexPath), exclusions}
}

// return a string slice for the results for a search string m with a json result string
func (si *SearchIndex) FileResults(query string) []string {
	matches := si.Matcher.ClosestList(query, 20)
	matchStrings := make([]string, len(matches))
	for pos, value := range matches {
		matchStrings[pos] = value.Data["fullPath"]
	}
	return matchStrings
}

func (si *SearchIndex) CodeResults(query string) []string {
	pat := "(?m)" + query
	re, err := csregexp.Compile(pat)
	if err != nil {
		log.Fatal(err)
	}
	codesearchQuery := index.RegexpQuery(re.Syntax)
	post := si.CodeIndex.PostingQuery(codesearchQuery)
	matchStrings := make([]string, len(post))
	for pos, fileid := range post {
		name := si.CodeIndex.Name(fileid)
		matchStrings[pos] = name
	}
	return matchStrings
}
