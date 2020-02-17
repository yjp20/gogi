package client

//go:generate echo "Generating Sass..."
//go:generate sass src/scss/main.scss src/css/main.css
//go:generate echo "Generating VFS Static File..."
//go:generate go run assets_generate.go

import (
	"io/ioutil"
	"log"
)

func getFile(s string) []byte {
	f, err := Assets.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func addFile(m map[string][]byte, s string) {
	m[s] = getFile(s)
}

func Templates() map[string][]byte {
	m := make(map[string][]byte)
	addFile(m, "ui/index.go")
	addFile(m, "ui/dispatcher.go")
	addFile(m, "ui/page_login.go")
	addFile(m, "ui/go.mod")
	addFile(m, "ui/go.sum")
	return m
}
