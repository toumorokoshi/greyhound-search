package greyhound

import "fmt"
import "log"
import "regexp"
import "encoding/json"

func NewGreyhoundSearch() *GreyhoundSearch {
	return &GreyhoundSearch{make(map[string]*SearchIndex)}
}

type GreyhoundSearch struct {
	Projects map[string]*SearchIndex
}

func (gs *GreyhoundSearch) AddProject(name, path string, exclusions []string) {
	regexExclusions := make([]*regexp.Regexp, len(exclusions), len(exclusions))
	for p, v := range exclusions {
		var err error
		regexExclusions[p], err = regexp.Compile(v)
		if err != nil {
			log.Print(err)
		}
	}
	gs.Projects[name] = NewSearchIndex(path, regexExclusions)
}

func (gs *GreyhoundSearch) PrintProjects() {
	for k, _ := range gs.Projects {
		fmt.Println(k)
	}
}


// return a search result for a projectName query
func (gs *GreyhoundSearch) Search(projectName, query string) string {
	_, hasKey := gs.Projects[projectName]
	var out_json []byte
	if hasKey {
		out_json, _ = json.Marshal(gs.Projects[projectName].Results(query))
	} else {
		out_json, _ = json.Marshal([]string{"no results found"})
	}
	return string(out_json)
}
