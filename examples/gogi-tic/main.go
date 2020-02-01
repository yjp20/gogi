package main

import (
	"github.com/yjp20/gogi"

	"github.com/gorilla/sessions"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
		"tic-tac-toe",
		&UserModel{},
		&RoomModel{},
		&ManagerModel{},

		gogi.WithDBProvider("sqlite3", "tmp.db"),
		gogi.WithSessionStore(sessions.NewCookieStore([]byte("WHAT"))),
		gogi.WithAuthMethod(&gogi.GuestAuth{}),
		gogi.WithAuthMethod(&gogi.LoginAuth{}),
	)
	g.Listen(":3000")
}
