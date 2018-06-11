package messaging

import (
	"auth"
	"math/rand"
	"user"
)

// Deploy listens from a channel and deployes the message to every client
func (channel *Channel) Deploy() {

	for {
		msg := <-channel.Tunnel
		for _, mapValues := range channel.UserSets {
			for client, socket := range mapValues {
				err := socket.WriteJSON(msg)

				if err != nil {
					socket.Close()
					delete(mapValues, client)
				}
			}
		}
	}
}

// FindRandomInSet finds a random user in a given set
func (channel *Channel) FindRandomInSet(setName string) *user.User {
	index := rand.Intn(len(channel.UserSets[setName]))
	c := 0
	for user := range channel.UserSets[setName] {
		if index == c {
			return user
		}
		c++
	}
	return nil
}

// HandelUser handels the user messages
func (channel *Channel) HandelUser(user *user.User) {

	moderator := channel.FindRandomInSet("M")

	for {
		var userMsg UserMessage
		var msg Message

		err := channel.UserSets["U"][user].ReadJSON(&userMsg)
		if err != nil {
			delete(channel.UserSets["U"], user)
			break
		}

		if _, ok := channel.UserSets["M"][moderator]; ok {
			msg = Message{user, moderator, userMsg.Text, "user_query"}
			channel.Tunnel <- msg
		} else {
			moderator = channel.FindRandomInSet("M")
			msg = Message{user, user, ModeratorUnavailableMsg, "self_message"}
			channel.Tunnel <- msg
		}

	}
}

// HandelModerator handels the moderator messages
func (channel *Channel) HandelModerator(user *user.User) {

	for {
		var moderatorMsg ModeratorMessage
		var msg Message

		err := channel.UserSets["M"][user].ReadJSON(&moderatorMsg)
		if err != nil {

			delete(channel.UserSets["M"], user)
			break
		}

		msg = Message{user, moderatorMsg.To, moderatorMsg.Text, "moderator_message"}

		channel.Tunnel <- msg

	}
}

// Wrapper wraps the services with clients and info to give handlerfunc
func Wrapper(channel Channel, fn func(Service)) func(auth.Service) {
	return func(sAuth auth.Service) {
		s := Service{}
		s.Service = sAuth
		s.Channel = channel
		fn(s)
	}
}
