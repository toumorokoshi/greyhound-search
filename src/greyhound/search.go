package greyhound

import "fmt"
import "log"
import "encoding/json"

func NewGreyhoundSearch() *GreyhoundSearch {
	return &GreyhoundSearch{make(map[string]*SearchIndex)}
}

type GreyhoundSearch struct {
	Projects map[string]*SearchIndex
}

func (gs *GreyhoundSearch) AddProject(path string) {
	gs.Projects[path] = NewSearchIndex(path)
}

func (gs *GreyhoundSearch) PrintProjects() {
	for k, v := range gs.Projects {
		fmt.Println(k)
		fmt.Println(v)
	}
}


// return a search result for a projectName query
func (gs *GreyhoundSearch) Search(projectName, query string) string {
	log.Print(projectName)
	_, hasKey := gs.Projects[projectName]
	var out_json []byte
	if hasKey {
		out_json, _ = json.Marshal(gs.Projects[projectName].Results(query))
	} else {
		out_json, _ = json.Marshal([]string{"no results found"})
	}
	return string(out_json)
}
