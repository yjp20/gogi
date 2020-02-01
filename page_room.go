package gogi

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var roomTemplateParsed *template.Template
var roomTemplate = `<!doctype html>
<html>
<head>
	<title> {{.Context.Name}} </title>
	<meta name="description" content="{{.Context.Description}}">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.8.0/css/bulma.min.css">
	<style> .card { margin-bottom: 1em; border: 1px solid #ccc; border-radius: 5px; box-shadow: none; } </style>
</head>
<body>
	<section class="hero has-background-light is-fullheight">
		<div class="hero-body">
			<div class="container">
				<div class="card">
					<div class="card-content">
						<div class="columns">
							<div class="column is-8 is-9-widescreen">
								<h1 class="title"> {{.Room.GetName}} </h1>
								{{range .Room.GetUsers}}
									<p> {{.}} </p>
									<p> {{.GetNick}} </p>
								{{end}}
							</div>
							<div class="column is-4 is-3-widescreen">
								<div class="field">
									<form action="{{.Context.Prefix}}/room/new" method="POST">
										<a class="button is-fullwidth" href=""> Private room </a>
									</form>
								</div>
								<div class="field">
									<a class="button is-fullwidth is-danger" href="{{.Context.Prefix}}/room/id/{{.Room.GetShortID}}/leave"> leave </a>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</section>
</body>
</html>`

func init() {
	var err error
	roomTemplateParsed = template.New("").Funcs(templateFunctions)
	_, err = roomTemplateParsed.Parse(roomTemplate)

	if err != nil {
		fmt.Printf("%s", err)
	}
}

func (g *Game) roomNewHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		s, _ := g.Context.Store.Get(r, "session")
		u, ok := s.Values["self"].(uint)
		if !ok {
			http.Redirect(w, r, g.Context.Prefix+"/", 302)
			return
		}
		user := g.Context.NewUser()
		err := g.Context.DB.First(user, u).Error
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		room := g.Context.NewRoom()
		room.Init()
		room.AddUser(user)
		g.Context.Rooms[room.GetShortID()] = room
		http.Redirect(w, r, g.Context.Prefix+"/room/id/"+room.GetShortID(), 302)
	}
}

func (g *Game) roomHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		s, _ := g.Context.Store.Get(r, "session")
		u, ok := s.Values["self"].(uint)
		if !ok {
			http.Redirect(w, r, g.Context.Prefix+"/", 302)
			return
		}
		user := g.Context.NewUser()
		err := g.Context.DB.First(user, u).Error
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		room := g.Context.Rooms[ps.ByName("id")]
		err = roomTemplateParsed.Execute(w, &RenderData{
			Context: g.Context,
			Self:    user,
			Room:    room,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Game) roomIdLeaveHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	}
}
