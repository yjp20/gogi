package gogi

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var gameTemplateParsed *template.Template
var gameTemplate = `<!doctype html>
<html>
<head>
	<title> </title>
	<link rel="stylesheet" href="">
</head>
<body>
</body>
</html>
`

func init() {
	var err error
	gameTemplateParsed = template.New("")
	_, err = gameTemplateParsed.Parse(gameTemplate)

	if err != nil {
		fmt.Printf("%s", err)
	}
}

func (g *Game) gameHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	}
}
