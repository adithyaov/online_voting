package messaging

import (
	"user"
	"net/http"
	"github.com/gorilla/websocket"
	"auth"
)

func HandleConnectionUser(w http.ResponseWriter, r *http.Request,
					      clients map[*user.User]*websocket.Conn,
					      info map[string]int, ch chan Message) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	defer ws.Close()

	// read the token and populate the user
	// token := r.Header["token"][0]
	// googleToken, err := auth.ParseToken(token)

	user := user.User{}
	clients[&user] = ws

	handelUser(clients, info, &user, ch)
	

}


func HandleConnectionModerator(w http.ResponseWriter, r *http.Request,
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
		http.Error(w, "You are not a moderator.", 400)
	}
	

}





