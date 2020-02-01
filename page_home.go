package gogi

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var homeTemplateParsed *template.Template
var homeTemplate = `<!doctype html>
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
				<div class="columns">
					<div class="column is-6 is-offset-3 is-4-widescreen is-offset-4-widescreen">
						<h1 class="title is-size-1"> {{.Context.Name}} </h1>
						{{range .Context.AuthMethods}}
							<div class="card">
								{{exec .Template $}}
							</div>
						{{end}}
						<p class="has-text-grey has-text-centered"> built with <a href="https://github.com/yjp20/gogi">gogi</a> </p>
					</div>
				</div>
			</div>
		</div>
	</section>
</body>
</html>
`

func init() {
	var err error
	homeTemplateParsed = template.New("").Funcs(templateFunctions)
	_, err = homeTemplateParsed.Parse(homeTemplate)

	if err != nil {
		fmt.Printf("%s", err)
	}
}

func (g *Game) homeHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		s, _ := g.Context.Store.Get(r, "session")

		u, ok := s.Values["self"].(uint)
		log.Printf("%v %d", u)
		if ok {
			http.Redirect(w, r, g.Context.Prefix+"/rooms", 302)
			return
		}
		err := homeTemplateParsed.Execute(w, &RenderData{Context: g.Context})
		if err != nil {
			log.Fatal(err)
		}
	}
}
