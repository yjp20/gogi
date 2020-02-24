package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

type RoomsView struct {
	vecty.Core
}

func (r *RoomsView) Render() vecty.ComponentOrHTML {
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
							vecty.Class("card-body"),
						),
						elem.Heading1(
							vecty.Markup(
								vecty.Class("title"),
							),
							vecty.Text("Rooms"),
						),
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
				r.renderControls(),
			),
		),
	)
}

func (r *RoomsView) renderControls() vecty.ComponentOrHTML {
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
				elem.Button(
					vecty.Markup(
						vecty.Class("button"),
						vecty.Property("type", "button"),
					),
					vecty.Text("Find Random"),
				),
			),
		),
	)
}

type RoomsList struct {
	vecty.Core
	Rooms []Room
}

type Room struct {
	Name   string
	Public bool
	Min    int
	Max    int
	Leader string
}

func (rl *RoomsList) Render() vecty.ComponentOrHTML {
	rml := make([]vecty.Component, len(rl.Rooms))
	for i, room := range rl.Rooms {
		rml[i] = &RoomListing{Room: room}
	}

	return elem.Div(
		vecty.Markup(
			vecty.Class("roomslist-wrapper"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("roomslist"),
			),
			rml...,
		),
	)
}

type RoomListing struct {
	vecty.Core
	Room Room
}

type PrivateRoomButton struct {
	vecty.Core
}

func (p *PrivateRoomButton) onSubmit(e *vecty.Event) {

}

func (p *PrivateRoomButton) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("field"),
		),
		elem.Button(
			vecty.Markup(
				vecty.Class("button"),
				vecty.Property("type", "button"),
				event.Click(p.onSubmit),
			),
			vecty.Text("New Private Room"),
		),
	)
}
