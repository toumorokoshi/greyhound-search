package main

import "io"
import "log"
import "os"
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
	gs.HandleGreyhoundSearchSocket(ws)
}

func handleIndexPage(w http.ResponseWriter, req *http.Request) {
	fi, err := os.Open("index.html")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		if n2, err := w.Write(buf[:n]); err != nil {
			panic(err)
		} else if n2 != n {
			panic("error in writing")
		}
	}
	//io.WriteString(w, os.Open("index.html"))
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
