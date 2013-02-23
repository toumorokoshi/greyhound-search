package main

import "greyhound"

func main() {
	gs := &greyhound.GreyhoundSearch{make([]string, 0)}
	gs.AddProject("test");
	gs.PrintProjects()
}
