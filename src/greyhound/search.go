package greyhound

import "fmt"
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
	_, hasKey := gs.Projects[projectName]
	if hasKey {
		result, _ := json.Marshal(gs.Projects[projectName].Results(query))
		return string(result)
	}
	return "nothing found!"
}
