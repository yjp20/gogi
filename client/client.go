package client

//go:generate go run assets_generate.go

import (
	"io/ioutil"
	"log"
)

var WasmLoader = getFile("wasm_exec.js")

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
	addFile(m, "index.go")
	return m
}
