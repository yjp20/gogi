package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type LoginView struct {
	vecty.Core
}

func (l *LoginView) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Div(
			vecty.Markup(
				vecty.Class("columns"),
				vecty.Class("is-tablet"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("column"),
					vecty.Class("is-4-tablet"),
					vecty.Class("is-3-desktop"),
					vecty.Class("is-2-widescreen"),
					vecty.Class("is-offset-2-desktop"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("card"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("card-body"),
						),
						&GuestLogin{},
						&DBLogin{},
					),
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("column"),
					vecty.Class("is-8-tablet"),
					vecty.Class("is-5-desktop"),
					vecty.Class("is-6-widescreen"),
				),
				elem.Div(
					vecty.Markup(),
					elem.Heading1(
						vecty.Markup(
							vecty.Class("title"),
						),
						vecty.Text("tic-tac-toe"),
					),
					elem.Paragraph(
						vecty.Markup(
							vecty.Class("paragraph"),
						),
					),
				),
			),
		),
	)
}

type GuestLogin struct {
	vecty.Core
}

func (g *GuestLogin) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Div(
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
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("field"),
			),
			elem.Button(
				vecty.Markup(
					vecty.Class("button"),
					vecty.Property("type", "button"),
				),
				vecty.Text("Guest Login"),
			),
		),
	)
}

type DBLogin struct {
	vecty.Core
}

func (d *DBLogin) Render() vecty.ComponentOrHTML {
	return elem.Div()
}
