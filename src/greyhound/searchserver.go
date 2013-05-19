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

func (gs *GreyhoundSearch) HandleGreyhoundSearch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	project, hasProject := r.Form["project"]
  query, hasQuery := r.Form["query"]
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
func (gs *GreyhoundSearch) HandleGreyhoundSearchSocket(ws *websocket.Conn) {
	for {
		var msg socketMessage
		err := websocket.JSON.Receive(ws, &msg)
		log.Println("raw message: ",  msg)
		if err != nil { 
			fmt.Println(err)
			break
		}
		if strings.EqualFold(msg.Action, "query") {
			gs.socketQuery(ws, msg.QueryData)
		}
	}
}

func (gs *GreyhoundSearch) socketQuery(ws *websocket.Conn, qd queryData) {
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


