package main

import (
	"github.com/yjp20/godux"
)

type Actions interface{}

type Error struct {
	Message string
}

type Login struct {
	Token    string
	LoggedIn bool
}

type Store struct {
	godux.Store
	Token    string
	LoggedIn bool
}

func (s *Store) OnAction(action interface{}) {
	switch a := action.(type) {
	case *Login:
		s.Token = a.Token
		s.LoggedIn = a.LoggedIn
	}
	s.UpdateComponents()
}
