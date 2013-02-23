package main

import "net/http"
import "greyhound"

var gs = greyhound.NewGreyhoundSearch()

func handler(w http.ResponseWriter, r *http.Request) {
	greyhound.HandleGreyhoundSearch(w, r, gs)
}

func main() {
	gs.AddProject("/tmp/");
	gs.PrintProjects()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
