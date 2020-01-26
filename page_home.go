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
	<title> {{.Name}} </title>
	<meta name="description" content="{{ .Context.Description }}">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.8.0/css/bulma.min.css">
</head>
<body>
	<section class="hero has-background-light is-fullheight">
		<div class="hero-body">
			<div class="container">
				<div class="columns">
					<div class="column is-6 is-offset-3 is-4-widescreen is-offset-4-widescreen">
						<style> .card { margin-bottom: 1em; border: 1px solid #ccc; border-radius: 5px; box-shadow: none; } </style>
						<div class="card">
							<div class="card-content has-background-dark">
								<h1 class="has-text-white title"> {{.Name}} </h1>
							</div>
						</div>
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
		err := homeTemplateParsed.Execute(w, g)
		if err != nil {
			log.Fatal(err)
		}
	}
}
