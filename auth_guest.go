package gogi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

const guestAuthClient = `package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

type GuestAuth struct {
	vecty.Core
}

func (g *GuestAuth) onSubmit(e *vecty.Event) {
	go func() {
		resp, err := http.Post("{{.Prefix}}/auth/guest", "application/json", nil)
		if err != nil {
			d.Dispatch(&Error{
				Message: "Error getting login response",
			})
			log.Printf("Error getting login response: %v", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			d.Dispatch(&Error{
				Message: "Error reading login response",
			})
			log.Printf("Error reading login response: %v", err)
			return
		}
		m := make(map[string]string)
		err = json.Unmarshal(body, &m)
		if err != nil {
			d.Dispatch(&Error{
				Message: "Error unmarshalling login response",
			})
			log.Printf("Error unmarshalling login response: %v", err)
			return
		}
		d.Dispatch(&Login{
			Token: m["token"],
			LoggedIn: true,
		})
	}()
}

func (g *GuestAuth) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("card"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("card-body"),
			),
			g.renderNickField(),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("card-body"),
				vecty.Class("has-bg-lightblue"),
			),
			g.renderSubmitField(),
		),
	)
}

func (g *GuestAuth) renderNickField() vecty.ComponentOrHTML{
	return elem.Div(
		vecty.Markup(
			vecty.Class("field"),
		),
		elem.Paragraph(
			vecty.Markup(
				vecty.Class("label"),
			),
			vecty.Text("Nickname"),
		),
		elem.Input(
			vecty.Markup(
				vecty.Property("type", "text"),
				vecty.Property("name", "nick"),
			),
		),
	)
}

func (g *GuestAuth) renderSubmitField() vecty.ComponentOrHTML{
	return elem.Div(
		vecty.Markup(
			vecty.Class("field"),
		),
		elem.Button(
			vecty.Markup(
				vecty.Class("button"),
				vecty.Property("type", "button"),
				event.Click(g.onSubmit),
			),
			vecty.Text("Guest Login"),
		),
	)
}`

type GuestAuth struct{}

func (a *GuestAuth) Name() string {
	return "GuestAuth"
}

func (a *GuestAuth) Client() (string, string) {
	return "ui/auth_guest.go", guestAuthClient
}

func (a *GuestAuth) Handlers(c *Context) []AuthMethodHandler {
	guestHandler := AuthMethodHandler{
		"POST", "/auth/guest", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			nick := r.Form.Get("nick")
			u := c.NewUser()
			u.SetNick(nick)
			u.SetShortID()
			u.SetTemp(true)
			err = c.DB.Save(u).Error
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"self_id": u.GetID(),
				"nbf":     time.Now().Unix(),
			})
			tokenString, err := token.SignedString(c.Secret)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"token": tokenString,
			})
		},
	}
	return []AuthMethodHandler{guestHandler}
}
