package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/yjp20/godux"
)

var (
	d         godux.Dispatcher
	s         Store
	userToken string
)

func main() {
	s.Init()
	d.Init()
	d.Register(s.OnAction)
	vecty.RenderBody(s.Connect(&PageView{}))
}

type PageView struct {
	vecty.Core
	LoggedIn bool
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
				vecty.If(!p.LoggedIn, s.Connect(&LoginView{})),
				vecty.If(p.LoggedIn, &LoggedInView{}),
			),
		),
	)
}

func (p *PageView) Connect() map[interface{}]interface{} {
	return map[interface{}]interface{}{
		&p.LoggedIn: &s.LoggedIn,
	}
}

type LoggedInView struct {
	vecty.Core
	inRoom bool
}

func (l *LoggedInView) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&RoomsView{},
	)
}
