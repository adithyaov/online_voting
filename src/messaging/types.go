package messaging

import (
	"user"
)

type Message struct {
	From *user.User `json:"from"`
	To *user.User   `json:"to"`
	Text string 	`json:"text"`
}

type UserMessage struct {
	Text string `json:"text"`
}

type ModeratorMessage struct {
	To *user.User `json:"to"`
	Text string   `json:"text"`
}

