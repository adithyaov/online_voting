package messaging

import (
	c "common"
	"math/rand"
	"net/http"
	"user"

	"github.com/gorilla/websocket"
)

// DeployFromChannel listens from a channel and deployes the message to every client
func DeployFromChannel(clients map[*user.User]*websocket.Conn,
	info *c.ThreadSafeType, ch chan Message) {
	for {
		msg := <-ch
		for client, socket := range clients {
			err := socket.WriteJSON(msg)
			if err != nil {
				func() {
					info.Mutex.Lock()
					defer info.Mutex.Unlock()
					info.Data.(map[string]int)["num_moderators"]--
					socket.Close()
					delete(clients, client)
				}()
			}
		}
	}
}

func findModerator(clients map[*user.User]*websocket.Conn,
	info *c.ThreadSafeType) *user.User {

	info.Mutex.Lock()
	defer info.Mutex.Unlock()
	skip := rand.Intn(info.Data.(map[string]int)["num_moderators"])

	for client := range clients {
		if client.RoleCode == "M" {
			if skip == 0 {
				return client
			}
			skip--
		}
	}
	return nil
}

func handelUser(clients map[*user.User]*websocket.Conn,
	info *c.ThreadSafeType, user *user.User, ch chan Message) {

	moderator := findModerator(clients, info)

	for {
		var userMsg UserMessage
		var msg Message

		err := clients[user].ReadJSON(&userMsg)
		if err != nil {

			delete(clients, user)
			break
		}

		if _, ok := clients[moderator]; ok {
			msg = Message{user, moderator, userMsg.Text}
			ch <- msg
		} else {
			moderator = findModerator(clients, info)
			msg = Message{user, user, ModeratorUnavailableMsg}
			ch <- msg
		}

	}
}

func handelModerator(clients map[*user.User]*websocket.Conn,
	info *c.ThreadSafeType, user *user.User, ch chan Message) {
	info.Mutex.Lock()
	info.Data.(map[string]int)["num_moderators"]++
	info.Mutex.Unlock()

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

// Wrapper wraps the services with clients and info to give handlerfunc
func Wrapper(clients map[*user.User]*websocket.Conn,
	info map[string]int, ch chan Message, fn MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, clients, info, ch)
	}
}
