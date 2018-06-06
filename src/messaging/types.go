package messaging

import (
	"auth"
	"net/http"
	"user"

	"github.com/gorilla/websocket"
)

// Message is the basic message type
type Message struct {
	From *user.User `json:"from"`
	To   *user.User `json:"to"`
	Text string     `json:"text"`
	Type string     `json:"type"`
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
	map[*user.User]*websocket.Conn, map[string]int, chan Message)

// Channel is the binded version of clients and channel
type Channel struct {
	UserSets map[string]map[*user.User]*websocket.Conn
	Tunnel   chan Message
}

// Service is the type the API's take as input in this package
type Service struct {
	auth.Service
	Channel
}
