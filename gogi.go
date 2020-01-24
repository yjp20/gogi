package gogi

import (
	"net/http"
)

type Game struct {
	Name  string
	Rooms []interface{}
}

func (g *Game) RegisterTemplate(name string) error {

}
