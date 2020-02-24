package gogi

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const loginAuthClient = `package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

type LoginAuth struct {
	vecty.Core

	isSignup bool
}

func (d *LoginAuth) toSignup(e *vecty.Event) {
	d.isSignup = true
	vecty.Rerender(d)
}

func (d *LoginAuth) toLogin(e *vecty.Event) {
	d.isSignup = false
	vecty.Rerender(d)
}

func (d *LoginAuth) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("card"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("card-body"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("field"),
				),
				elem.Paragraph(
					vecty.Markup(
						vecty.Class("label"),
					),
					vecty.Text("Username"),
				),
				elem.Input(
					vecty.Markup(
						vecty.Property("type", "text"),
						vecty.Property("name", "username"),
					),
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("field"),
				),
				elem.Paragraph(
					vecty.Markup(
						vecty.Class("label"),
					),
					vecty.Text("Password"),
				),
				elem.Input(
					vecty.Markup(
						vecty.Property("type", "password"),
						vecty.Property("name", "password"),
					),
				),
			),
			vecty.If(d.isSignup,
				elem.Div(
					vecty.Markup(
						vecty.Class("field"),
					),
					elem.Paragraph(
						vecty.Markup(
							vecty.Class("label"),
						),
						vecty.Text("Password Confirm"),
					),
					elem.Input(
						vecty.Markup(
							vecty.Property("type", "password"),
							vecty.Property("name", "password_confirm"),
						),
					),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("field"),
					),
					elem.Paragraph(
						vecty.Markup(
							vecty.Class("label"),
						),
						vecty.Text("Email"),
					),
					elem.Input(
						vecty.Markup(
							vecty.Property("type", "email"),
							vecty.Property("name", "email"),
						),
					),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("card-body"),
				vecty.MarkupIf(!d.isSignup, vecty.Class("has-bg-lightred")),
				vecty.MarkupIf(d.isSignup, vecty.Class("has-bg-lightgreen")),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("field"),
				),
				vecty.If(d.isSignup,
					elem.Button(
						vecty.Markup(
							vecty.Class("button"),
							vecty.Property("type", "button"),
						),
						vecty.Text("Signup"),
					),
					elem.Button(
						vecty.Markup(
							vecty.Class("button"),
							vecty.Class("is-link"),
							event.Click(d.toLogin),
						),
						vecty.Text("Login"),
					),
				),
				vecty.If(!d.isSignup,
					elem.Button(
						vecty.Markup(
							vecty.Class("button"),
							vecty.Property("type", "button"),
						),
						vecty.Text("Login"),
					),
					elem.Button(
						vecty.Markup(
							vecty.Class("button"),
							vecty.Class("is-link"),
							event.Click(d.toSignup),
						),
						vecty.Text("Signup"),
					),
				),
			),
		),
	)
}`

type LoginAuth struct{}

func (a *LoginAuth) Name() string {
	return "LoginAuth"
}

func (a *LoginAuth) Client() (string, string) {
	return "ui/auth_login.go", loginAuthClient
}

func (a *LoginAuth) Handlers(c *Context) []AuthMethodHandler {
	loginHandler := AuthMethodHandler{
		"POST", "/auth/login", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		},
	}
	return []AuthMethodHandler{loginHandler}
}
