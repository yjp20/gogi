package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type LoginView struct {
	vecty.Core
}

func (l *LoginView) Connect() map[interface{}]interface{} {
	return map[interface{}]interface{}{}
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
					vecty.Class("is-8-tablet"),
					vecty.Class("is-9-desktop"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("card"),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("card-header"),
						),
						l.renderInformation(),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("card-body"),
						),
						l.renderDescription(),
					),
				),
				elem.Paragraph(
					vecty.Markup(
						vecty.Class("paragraph"),
						vecty.Class("has-fg-grey"),
					),
					vecty.Text("Built with Gogi"),
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("column"),
					vecty.Class("is-4-tablet"),
					vecty.Class("is-3-desktop"),
				),
				// Authmethods {{newline}} {{range .AuthMethods}} &{{.Name}}{}, {{end}}
			),
		),
	)
}

func (l *LoginView) renderInformation() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Heading1(
			vecty.Markup(
				vecty.Class("title"),
			),
			vecty.Text("{{.Name}}"),
		),
		vecty.If(len("{{.Version}}") > 0,
			elem.Heading2(
				vecty.Markup(
					vecty.Class("subtitle"),
				),
				vecty.Text("{{.Version}}"),
			),
		),
	)
}

func (l *LoginView) renderDescription() vecty.ComponentOrHTML {
	return elem.Paragraph(
		vecty.Markup(
			vecty.Class("paragraph"),
		),
		vecty.Text("{{.Description}}"),
	)
}
