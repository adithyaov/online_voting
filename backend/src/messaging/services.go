package messaging

import (
	"user"
)

// HandleConnectionUser is a service to handel User connections
func HandleConnectionUser(s Service) {

	ws, err := upgrader.Upgrade(s.Writer, s.Request, nil)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	defer ws.Close()

	user := user.User{}
	user.FromToken(s.Token)

	s.UserSets["U"][&user] = ws

	s.HandelUser(&user)

}

// HandleConnectionModerator is a service to handel Moderator connections
func HandleConnectionModerator(s Service) {

	ws, err := upgrader.Upgrade(s.Writer, s.Request, nil)
	if err != nil {
		s.Tell(err.Error(), 400)
		return
	}

	defer ws.Close()

	user := user.User{}
	user.FromToken(s.Token)

	s.UserSets["M"][&user] = ws

	s.HandelUser(&user)
}
