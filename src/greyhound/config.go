package greyhound

import "json"

type GreyhoundProjectConfig struct {
	Root string
	Exclusions []string
}

type GreyhoundConfig struct {
	Projects []GreyhoundProjectConfig
}

func (gs *GreyhoundSearch) LoadFromConfig(path string) {
}

func (gs *GreyhoundSearch) LoadFromConfigString(path string) {
}
