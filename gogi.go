package gogi

import (
	"encoding/gob"
	"log"
	"net/http"
	"reflect"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	// "github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
)

type Game struct {
	Context *Context
	Name    string
	Rooms   []interface{}
}

type Context struct {
	Name        string
	AuthMethods []AuthMethod
	Middlewares []func(http.Handler) http.Handler
	Prefix      string
	Description string
	DB          *gorm.DB
	Store       sessions.Store
	Rooms       map[string]Room

	UserModel    User
	RoomModel    Room
	ManagerModel Manager
}

func (c *Context) Init() {
	c.Rooms = make(map[string]Room)
}

func (c *Context) NewUser() User {
	a := reflect.ValueOf(c.UserModel).Elem() // Gets the user supplied model
	u := reflect.New(a.Type()).Interface().(User)
	return u
}

func (c *Context) NewRoom() Room {
	a := reflect.ValueOf(c.RoomModel).Elem() // Gets the user supplied model
	r := reflect.New(a.Type()).Interface().(Room)
	return r
}

type Option func(*Context)

func WithSessionStore(s sessions.Store) Option {
	return func(c *Context) {
		c.Store = s
	}
}

func WithGzip() Option {
	return func(c *Context) {
		c.Middlewares = append(c.Middlewares, gziphandler.GzipHandler)
	}
}

func WithPrefix(prefix string) Option {
	return func(c *Context) {
		c.Prefix = prefix
	}
}

func WithDescription(desc string) Option {
	return func(c *Context) {
		c.Description = desc
	}
}

func NewGame(name string, user User, room Room, manager Manager, options ...Option) Game {
	c := &Context{}
	c.Init()
	c.Name = name
	c.UserModel = user
	c.RoomModel = room
	c.ManagerModel = manager

	for _, option := range options {
		option(c)
	}

	if c.DB == nil {
		log.Fatal("You must register a DB provider with gogi.WithDBProvider(...)")
	}

	c.DB.AutoMigrate(user)
	gob.Register(user)
	gob.Register(room)
	gob.Register(manager)

	g := Game{Context: c, Name: name}
	return g
}

func (g *Game) Listen(port string) error {
	r := httprouter.New()

	r.GET(g.Context.Prefix+"/", g.homeHandler())
	r.GET(g.Context.Prefix+"/wasm", g.homeWASMHandler())
	r.GET(g.Context.Prefix+"/wasm.js", g.homeWASMLoaderHandler())
	r.GET(g.Context.Prefix+"/room/new", g.roomNewHandler())
	r.GET(g.Context.Prefix+"/auth/logout", g.logoutHandler())

	for _, authMethod := range g.Context.AuthMethods {
		for _, h := range authMethod.Handlers(g.Context) {
			r.Handle(h.Method, h.Slug, h.Handler)
		}
	}

	m := http.Handler(r)
	for _, mw := range g.Context.Middlewares {
		m = mw(m)
	}

	return http.ListenAndServe(port, m)
}
