package gogi

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
	// "github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
)

type Game struct {
	Context Context
	Name    string
	Rooms   []interface{}
}

type Context struct {
	AuthMethods []AuthMethod
	Middlewares []func(http.Handler) http.Handler
	Prefix      string
	Description string
	DB          interface{}
	UserModel   User
}

type Option func(*Context)

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

func NewGame(name string, options ...Option) Game {
	c := Context{}
	for _, option := range options {
		option(&c)
	}
	g := Game{Context: c, Name: name}
	return g
}

func (g *Game) Listen(port string) error {
	r := httprouter.New()

	r.GET(g.Context.Prefix+"/", g.homeHandler())
	r.GET(g.Context.Prefix+"/game", g.gameHandler())
	r.GET(g.Context.Prefix+"/rooms", g.roomsHandler())

	for _, authMethod := range g.Context.AuthMethods {
		for _, h := range authMethod.Handlers(&g.Context) {
			r.Handle(h.Method, h.Slug, h.Handler)
		}
	}

	m := http.Handler(r)
	for _, mw := range g.Context.Middlewares {
		m = mw(m)
	}

	return http.ListenAndServe(port, m)
}
