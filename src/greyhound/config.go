package greyhound

import "encoding/json"
import "io/ioutil"
import "log"

type GreyhoundProjectConfig struct {
	Root string
	Exclusions []string
}

type GreyhoundConfig struct {
	Projects map[string]GreyhoundProjectConfig
}

func (gs *GreyhoundSearch) LoadFromConfig(path string) {
	var gc GreyhoundConfig
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Print("Could not read file! ", err)
	}
	err = json.Unmarshal(contents, &gc)
	if err != nil {
		log.Print("Could not unmarshal file! ", err)
	}
	if err == nil {
			gs.LoadFromConfigStruct(&gc)
	}
}

func (gs *GreyhoundSearch) LoadFromConfigStruct(gc *GreyhoundConfig) {
	for name, project := range gc.Projects {
		go gs.AddProject(name, project.Root, project.Exclusions)
	}
}
