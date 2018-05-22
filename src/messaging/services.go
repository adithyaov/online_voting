package messaging

import (
	"user"
	"net/http"
	"github.com/gorilla/websocket"
)

func HandleConnection(w http.ResponseWriter, r *http.Request,
					  clients map[*user.User]*websocket.Conn,
					  info map[string]int, ch chan Message) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	defer ws.Close()

	// read the token and populate the user
	user := user.User{}
	clients[&user] = ws

	// handel token and dependigly execute one of these
	handelUser(clients, info, &user, ch)
	handelModerator(clients, info, &user, ch)

}





