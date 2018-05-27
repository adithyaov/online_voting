package messaging

import (
	"user"
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
