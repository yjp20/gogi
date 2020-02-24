package main

import (
	"github.com/yjp20/gogi"
)

type UserModel struct {
	gogi.UserModel // Must extend user
}

type RoomModel struct {
	gogi.RoomModel

	Users map[string]*UserModel
	Black *UserModel
	White *UserModel
	Board [][]int
	Turn  bool
}

type ManagerModel struct {
	gogi.ManagerModel

	Rooms map[string]*RoomModel
}

func (s *ManagerModel) Events() {
	s.Register("start", func(d gogi.MessageModel) {
		rm := d.Room.(*RoomModel)
		rm.Board = make([][]int, 3)
		for i := 0; i < 3; i++ {
			rm.Board[i] = make([]int, 3)
		}
	})

	s.Register("move", func(d gogi.MessageModel) {
	})
}

func main() {
	g := gogi.NewGame(
		"Tic Tac Toe",
		&UserModel{},
		&RoomModel{},
		&ManagerModel{},

		gogi.WithAuthMethod(&gogi.GuestAuth{}),
		gogi.WithAuthMethod(&gogi.LoginAuth{}),
		gogi.WithDescription("Tic Tac Toe is a game that can be traced to Ancient Egypt, 1300 BCE, and is one of the most iconic games of the modern world. This particular implementation of Tic Tac Toe features the usage of the Gogi library, which makes it easy to make games like this without having to deal with the complexity of user login, room handling, match making."),
		gogi.WithVersion("b0.0.1"),
	)
	g.Listen(":3000")
}
