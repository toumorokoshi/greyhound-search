package main

import "log"
import "net/http"

import "code.google.com/p/go.net/websocket"

import "greyhound"

var gs = greyhound.NewGreyhoundSearch()
// a list of regex exclusions from the workspace
var baseExclusions = []string{
	".*\\.class",
	".*\\.pyc",
	"\\.keep",
	".*\\.key",
	"\\.rspec",
}

func handleQuery(w http.ResponseWriter, r *http.Request) {
	gs.HandleGreyhoundSearch(w, r)
}

func handlerSocket(ws *websocket.Conn) {
	log.Print("handling socket...")
	gs.HandleGreyhoundSearchSocket(ws)
}

func handleIndexPage(w http.ResponseWriter, req *http.Request) {
	log.Print("handling index...")
	greyhound.HandleFile(w, "./index.html")
}

func main() {
	log.Print("Loading config...")
	gs.LoadFromConfig("config.json")
	log.Print(gs.ListProjects())
	http.Handle("/socket", websocket.Handler(handlerSocket))
	http.Handle("/statics/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/query", handleQuery)
	http.HandleFunc("/", handleIndexPage)
	log.Print("Listening on port 8081...")
	http.ListenAndServe(":8081", nil)
}
