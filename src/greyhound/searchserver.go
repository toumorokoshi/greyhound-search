/*
Code for a greyhound search server

TBD
*/
package greyhound

import "fmt"
import "net/http"

func HandleGreyhoundSearch(w http.ResponseWriter, r *http.Request, gs *GreyhoundSearch) {
	r.ParseForm()
	project, hasProject := r.Form["project"]
  query, hasQuery := r.Form["query"]
	if hasProject && hasQuery {
		fmt.Fprintf(w, gs.Search(project[0], query[0]))
	} else {
		fmt.Fprintf(w, "no query found!")
	}
}
