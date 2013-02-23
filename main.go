package main

import (
	"fmt"
	"net/http"
	"greyhound"
)

var gs = greyhound.NewGreyhoundSearch()

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, gs.Search("/tmp", "test"))
}

func main() {
	gs.AddProject("/tmp/");
	gs.PrintProjects()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
