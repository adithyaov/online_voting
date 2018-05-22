package messaging

import (
	"github.com/gorilla/websocket"
	"math/rand"
	"user"
)
	
func DeployFromChannel(clients map[*websocket.Conn]bool, ch chan Message) {
	for {
		msg := <-ch
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {

				client.Close()
				delete(clients, client)
			}
		}
	}
}

func findModerator(clients map[*user.User]*websocket.Conn,
				   info map[string]int) *user.User {
	skip := rand.Intn(info["num_moderators"])

	for k, _ := range clients {
		if k.RoleCode == "M" {
			if skip == 0 {
				return k
			}
			skip -= 1
		}
	}
	return nil
}


func handelUser(clients map[*user.User]*websocket.Conn,
				info map[string]int, user *user.User, ch chan Message) {

	// read the token and populate the user

	moderator := findModerator(clients, info)

	for {
		var userMsg UserMessage
		var msg Message

		err := clients[user].ReadJSON(&userMsg)
		if err != nil {

			delete(clients, user)
			break
		}

		if _, ok := clients[moderator]; ok == false {
			moderator = findModerator(clients, info)
			msg = Message{user, user, ModeratorUnavailableMsg}
			ch <- msg
		} else {
			msg = Message{user, moderator, userMsg.Text}
			ch <- msg
		}

		
	}
}

func handelModerator(clients map[*user.User]*websocket.Conn,
				     info map[string]int, user *user.User, ch chan Message) {

	// read the token and populate the moderator

	for {
		var moderatorMsg ModeratorMessage
		var msg Message

		err := clients[user].ReadJSON(&moderatorMsg)
		if err != nil {

			delete(clients, user)
			break
		}

		msg = Message{user, moderatorMsg.To, moderatorMsg.Text}

		ch <- msg

		
	}
}







