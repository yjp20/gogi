package client

//go:generate echo "Generating Sass..."
//go:generate sass src/scss/main.scss src/css/main.css
//go:generate echo "Generating VFS Static File..."
//go:generate go run assets_generate.go

import (
	"io/ioutil"
	"log"
	"strings"
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
	for k, v := range Assets.(vfsgen۰FS) {
		switch v.(type) {
		case *vfsgen۰DirInfo:
			break
		default:
			if strings.HasPrefix(k, "/ui") {
				addFile(m, k)
			}
		}
	}
	return m
}
