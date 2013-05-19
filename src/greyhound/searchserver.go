/*
Code for a greyhound search server

TBD
*/
package greyhound

import "fmt"
import "log"
import "net/http"
import "code.google.com/p/go.net/websocket"

func (gs *GreyhoundSearch) HandleGreyhoundSearch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	action, hasAction := r.Form["action"]
	if !hasAction {
		fmt.Fprintf(w, "no action argument passed!")
	} else {
		var queryData map[string]string
		for k, v := range r.Form {
			queryData[k] = v[0]
		}
		msg := &Message{action[0], queryData}
		gs.PerformAction(msg)
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
		var msg Message
		err := websocket.JSON.Receive(ws, &msg)
		log.Println("raw message: ",  msg)
		if err != nil { 
			fmt.Println(err)
			break
		}
		_ = websocket.Message.Send(ws, gs.PerformAction(&msg))
	}
}

func (gs *GreyhoundSearch) PerformAction (m *Message) string {
	switch m.Action {
	case "query": 
		return gs.Search(m.QueryData["project"], m.QueryData["query"])
  case "list_projects":
		return gs.ListProjects()
	default:
		return fmt.Sprintf("%s is not a valid action", m.Action)
	}
	// this code path is never hit
	return ""
}

type Message struct {
	Action string
	QueryData map[string]string
}
