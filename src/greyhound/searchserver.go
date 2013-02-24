/*
Code for a greyhound search server

TBD
*/
package greyhound

import "fmt"
import "strings"
import "net/http"
import "encoding/json"
import "code.google.com/p/go.net/websocket"

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
		if err != nil { 
			fmt.Println(err)
			break
		}
		if strings.EqualFold(msg.Action, "query") {
			socketQuery(ws, msg.Data, gs)
		}
	}
}

func socketQuery(ws *websocket.Conn, raw_json []byte, gs *GreyhoundSearch) {
	var qd queryData
	_ = json.Unmarshal(raw_json, qd)
	_ = websocket.Message.Send(ws, gs.Search(qd.Project, qd.Query))
}

type socketMessage struct {
	Action string
	Data []byte
}

type queryData struct {
	Project string
	Query string
}


