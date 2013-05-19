/*
Code for a greyhound search server

TBD
*/
package greyhound

import "fmt"
import "log"
import "strings"
import "net/http"
import "code.google.com/p/go.net/websocket"

func HandleGreyhoundSearch(w http.ResponseWriter, r *http.Request, gs *GreyhoundSearch) {
	r.ParseForm()
	project, hasProject := r.Form["project"]
  query, hasQuery := r.Form["query"]
	log.Print(project)
	log.Print(query)
	if hasProject && hasQuery {
		fmt.Fprintf(w, gs.Search(project[0], query[0]))
	} else {
		fmt.Fprintf(w, "no query found!")
	}
}


/* handles greyhound-search's websocket actions.
effectively, greyhound messages are always sent as json. Specifically:
{ action: 'ACTION',
  data: { JSON_OBJECT }
}

each action has a struct to unmarshal json, and returns a series of values
 */

func HandleGreyhoundSearchSocket(ws *websocket.Conn, gs *GreyhoundSearch) {
	for {
		var msg socketMessage
		err := websocket.JSON.Receive(ws, &msg)
		log.Println("raw message: ",  msg)
		if err != nil { 
			fmt.Println(err)
			break
		}
		if strings.EqualFold(msg.Action, "query") {
			socketQuery(ws, msg.QueryData, gs)
		}
	}
}

func socketQuery(ws *websocket.Conn, qd queryData, gs *GreyhoundSearch) {
	log.Println(qd)
	_ = websocket.Message.Send(ws, gs.Search(qd.Project, qd.Query))
}

type socketMessage struct {
	Action string
	QueryData queryData
}

type queryData struct {
	Project string
	Query string
}


