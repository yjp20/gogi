package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"marwan.io/vecty-router"
)

var d Dispatcher

func main() {
	d = Dispatcher{}
	d.Init()
	vecty.RenderBody(&PageView{})
}

type PageView struct {
	vecty.Core
}

func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Class("section"),
				vecty.Class("has-centered"),
				vecty.Class("has-bg-light"),
				vecty.Class("is-fullheight"),
			),
			elem.Div(
				vecty.Markup(vecty.Class("container")),
				router.NewRoute("/", &LoginView{}, router.NewRouteOpts{ExactMatch: true}),
				router.NewRoute("/rooms", &LoginView{}, router.NewRouteOpts{}),
				router.NewRoute("/room/{id}", &LoginView{}, router.NewRouteOpts{}),
			),
		),
	)
}
