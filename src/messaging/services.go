package messaging

import (
	"user"
	"net/http"
	"github.com/gorilla/websocket"
	"auth"
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

	token := r.Header["token"][0]
	googleToken, err := auth.ParseToken(token)
	// handel token and dependigly execute one of these
	if googleToken.RoleCode == "M" {
		handelModerator(clients, info, &user, ch)
	} else {
		handelUser(clients, info, &user, ch)
	}
	

}





