package gogi

import (
	"html/template"
	"log"
	"net/http"

	"github.com/yjp20/gogi/client"

	"github.com/julienschmidt/httprouter"
)

var homeTemplateParsed *template.Template
var homeTemplate = `<!doctype html>
<html>
<head>
	<title> {{.Context.Name}} </title>
	<meta name="description" content="{{ .Context.Description }}">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="{{.Context.Prefix}}/static/css/main.css">
</head>
<body>
	<div id="router"></div>
	<script src="{{.Context.Prefix}}/static/js/wasm_exec.js"></script>
	<script>
		const go = new Go();
		WebAssembly.instantiateStreaming(fetch("{{.Context.Prefix}}/wasm"), go.importObject).then((result) => {
			go.run(result.instance);
		});
	</script>
</body>
</html>
`

var wasmBinary []byte

func init() {
	var err error
	homeTemplateParsed = template.New("").Funcs(templateFunctions)
	_, err = homeTemplateParsed.Parse(homeTemplate)
	if err != nil {
		log.Fatalf("%v", err)
	}
	wasmBinary, err = client.CompileIndex(struct{}{})
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func (g *Game) homeHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := homeTemplateParsed.Execute(w, &RenderData{
			Context: g.Context,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Game) homeWASMHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write(wasmBinary)
	}
}
