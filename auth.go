package gogi

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthMethodHandler struct {
	Method  string
	Slug    string
	Handler httprouter.Handle
}

type AuthMethod interface {
	Init(*Context) error
	Template() *template.Template
	Handlers(*Context) []AuthMethodHandler
}

func WithAuthMethod(am AuthMethod) Option {
	return func(c *Context) {
		err := am.Init(c)
		if err != nil {
			log.Fatal(err)
		}
		c.AuthMethods = append(c.AuthMethods, am)
	}
}

func (g *Game) logoutHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		s, _ := g.Context.Store.Get(r, "session")
		_, ok := s.Values["self"].(uint)
		if ok {
			delete(s.Values, "self")
		}
		err := s.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		http.Redirect(w, r, g.Context.Prefix+"/", 302)
	}
}
