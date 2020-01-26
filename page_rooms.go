package gogi

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var roomsTemplateParsed *template.Template
var roomsTemplate = `<!doctype html>
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
	roomsTemplateParsed = template.New("")
	_, err = roomsTemplateParsed.Parse(roomsTemplate)

	if err != nil {
		fmt.Printf("%s", err)
	}
}

func (g *Game) roomsHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	}
}
