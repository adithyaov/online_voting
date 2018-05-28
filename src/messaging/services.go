package messaging

import (
	"auth"
	"net/http"
	"user"

	"github.com/gorilla/websocket"
)

// HandleConnectionUser is a service to handel User connections
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
	token := r.Header["token"][0]
	googleToken, err := auth.ParseToken(token)

	user := user.User{}
	user.FromToken(googleToken)

	clients[&user] = ws

	handelUser(clients, info, &user, ch)

}

// HandleConnectionModerator is a service to handel Moderator connections
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
	token := r.Header["token"][0]
	googleToken, err := auth.ParseToken(token)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user := user.User{}
	user.FromToken(googleToken)

	clients[&user] = ws

	// handel token and dependigly execute one of these
	if user.RoleCode == "M" {
		handelModerator(clients, info, &user, ch)
	} else {
		http.Error(w, "You are not a moderator.", 400)
	}

}
