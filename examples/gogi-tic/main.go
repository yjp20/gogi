package main

import (
	"github.com/yjp20/gogi"
)

type StateModel struct {
	gogi.StateModel
}

func (s *StateModel) Events() {
	s.Register("init", func(d gogi.Message) {
		
	}

	s.Register("move", func(d gogi.Message) {
		d.
	}
}

type UserModel struct {
	gogi.UserModel // Must extend user
}

func main() {
	g := gogi.NewGame("tic-tac-toe",
		gogi.WithDBProvider("sqlite3", "/tmp/gorm.db"),
		gogi.WithAuthMethod(&gogi.GuestAuth{}),
		gogi.WithAuthMethod(&gogi.LoginAuth{}),
		gogi.WithUserModel(&UserModel{}),
	)
	g.Listen(":3000")
}
