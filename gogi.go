package gogi

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/yjp20/gogi/client"

	"github.com/julienschmidt/httprouter"
)

type Game struct {
	Context *Context
	Name    string
	Rooms   []interface{}
}

func NewGame(name string, user User, room Room, manager Manager, options ...Option) Game {
	c := &Context{}
	c.Name = name
	c.UserModel = user
	c.RoomModel = room
	c.ManagerModel = manager

	for _, option := range options {
		option(c)
	}

	c.Init()

	c.DB.AutoMigrate(user)
	gob.Register(user)
	gob.Register(room)
	gob.Register(manager)

	g := Game{Context: c, Name: name}
	g.Init()
	return g
}

func (g *Game) Init() {
	g.InitClient()
}

func (g *Game) Listen(port string) error {
	r := httprouter.New()

	r.GET(g.Context.Prefix+"/", g.clientHandler())
	r.GET(g.Context.Prefix+"/wasm", g.clientWASMHandler())
	r.POST(g.Context.Prefix+"/room/new", g.roomNewHandler())
	r.ServeFiles(g.Context.Prefix+"/static/*filepath", client.Assets)

	for _, authMethod := range g.Context.AuthMethods {
		for _, h := range authMethod.Handlers(g.Context) {
			r.Handle(h.Method, h.Slug, h.Handler)
		}
	}

	m := http.Handler(r)
	for _, mw := range g.Context.Middlewares {
		m = mw(m)
	}

	log.Println("Starting Webserver")
	return http.ListenAndServe(port, m)
}
