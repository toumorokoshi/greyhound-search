package greyhound

import "fmt"


type GreyhoundSearch struct {
	Projects []string
}

func (gs *GreyhoundSearch) AddProject(path string) {
	gs.Projects = append(gs.Projects, path)
}

func (gs *GreyhoundSearch) PrintProjects() {
	l := len(gs.Projects)
	for i := 0; i < l; i++ {
		fmt.Println(gs.Projects[i])
	}
}


