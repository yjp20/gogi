package gogi

import (
	"html/template"
	"log"
	"net/http"

	"github.com/yjp20/gogi/client"

	"github.com/julienschmidt/httprouter"
)

var wasmBinary []byte
var homeTemplateParsed *template.Template
var homeTemplate = `<!doctype html>
<html>
<head>
	<title> {{.Name}} </title>
	<meta name="description" content="{{.Description }}">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="{{.Prefix}}/static/css/main.css">
</head>
<body>
	<div id="router"></div>
	<script src="{{.Prefix}}/static/js/wasm_exec.js"></script>
	<script>
		const go = new Go();
		WebAssembly.instantiateStreaming(fetch("{{.Prefix}}/wasm"), go.importObject).then((result) => {
			go.run(result.instance);
		});
	</script>
</body>
</html>
`

func (g *Game) InitClient() {
	var err error
	homeTemplateParsed = template.New("")
	_, err = homeTemplateParsed.Parse(homeTemplate)
	if err != nil {
		log.Fatalf("%v", err)
	}
	authClientFiles := make(map[string][]byte)
	for _, am := range g.Context.AuthMethods {
		filename, script := am.Client()
		authClientFiles[filename] = []byte(script)
	}
	wasmBinary, err = client.CompileIndex(g.Context, authClientFiles)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func (g *Game) clientHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := homeTemplateParsed.Execute(w, g.Context)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Game) clientWASMHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write(wasmBinary)
	}
}
