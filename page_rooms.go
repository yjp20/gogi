package gogi

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var roomsTemplateParsed *template.Template
var roomsTemplate = `<!doctype html>
<html>
<head>
	<title> {{.Context.Name}} </title>
	<meta name="description" content="{{ .Context.Description }}">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.8.0/css/bulma.min.css">
	<style> .card { margin-bottom: 1em; border: 1px solid #ccc; border-radius: 5px; box-shadow: none; } </style>
</head>
<body>
	<section class="hero has-background-light is-fullheight">
		<div class="hero-body">
			<div class="container">
				<p>Hello, {{.Self.GetNick}}</p>
				<div class="columns">
					<div class="column is-8 is-9-widescreen">
						<div class="card">
							<div class="card-content">
								<h1 class="title"> Rooms </h1>
								{{range .Context.Rooms}}
									<div class="card">
										{{.}}
									</div>
								{{end}}
							</div>
						</div>
					</div>
					<div class="column is-4 is-3-widescreen">
						<div class="card">
							<div class="card-content">
								<div class="field">
									<a class="button is-fullwidth" href="{{.Context.Prefix}}/room/new"> Private room </a>
								</div>
								<div class="field">
									<a class="button is-fullwidth is-danger" href="{{ .Context.Prefix }}/auth/logout"> logout </a>
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
	roomsTemplateParsed = template.New("").Funcs(templateFunctions)
	_, err = roomsTemplateParsed.Parse(roomsTemplate)

	if err != nil {
		fmt.Printf("%s", err)
	}
}

func (g *Game) roomsHandler() httprouter.Handle {
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
			http.Redirect(w, r, g.Context.Prefix+"/", 302)
			return
		}
		log.Printf("%+v", user)
		err = roomsTemplateParsed.Execute(w, &RenderData{
			Context: g.Context,
			Self:    user,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
