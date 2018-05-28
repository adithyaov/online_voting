package messaging

import (
	"net/http"
	"user"

	"github.com/gorilla/websocket"
)

// Message is the basic message type
type Message struct {
	From *user.User `json:"from"`
	To   *user.User `json:"to"`
	Text string     `json:"text"`
}

// UserMessage is the type a user sends, only to server
type UserMessage struct {
	Text string `json:"text"`
}

// ModeratorMessage is the type moderator sends, can direct his messages
type ModeratorMessage struct {
	To   *user.User `json:"to"`
	Text string     `json:"text"`
}

// MessageService is the expected type of function for any services related to websockets
type MessageService func(http.ResponseWriter, *http.Request,
	map[*user.User]*websocket.Conn, map[string]int,
	chan Message, websocket.Upgrader)
