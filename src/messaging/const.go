package messaging

import (
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


var ModeratorUnavailableMsg = "We could not find a moderator for you, Please get back after 5 minutes."

