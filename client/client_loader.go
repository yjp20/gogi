package client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"text/template"
	"time"
)

const (
	maxCompileTime = 30 * time.Second
	buildPkgArg    = "."
	goBin          = "go"
)

func findGopath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		usr, err := user.Current()
		if err != nil {
			return ""
		}
		usrgo := filepath.Join(usr.HomeDir, "go")
		return usrgo
	}
	return gopath
}

func RenderTemplates(templates map[string][]byte, data interface{}) (map[string][]byte, error) {
	rendered := make(map[string][]byte)
	for key, val := range templates {
		b := bytes.NewBuffer([]byte{})
		tmpl, err := template.New("").Parse(string(val))
		if err != nil {
			return rendered, err
		}
		err = tmpl.Execute(b, data)
		if err != nil {
			return rendered, err
		}
		rendered[key] = b.Bytes()
	}
	return rendered, nil
}

func CompileIndex(data interface{}) ([]byte, error) {
	rendered, err := RenderTemplates(Templates(), data)
	if err != nil {
		return []byte(""), err
	}
	return CompileWasmFrom(rendered)
}

func CompileWasmFrom(source map[string][]byte) ([]byte, error) {
	tmpDir, err := ioutil.TempDir("", "gogi-frontend-compile")
	if err != nil {
		return []byte{}, fmt.Errorf("error creating temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	err = os.Mkdir(filepath.Join(tmpDir, "ui"), 0755)
	if err != nil {
		return []byte{}, fmt.Errorf("error creating ui folder in temp directory: %v", err)
	}

	ctx := context.TODO()
	buildCtx, cancel := context.WithTimeout(ctx, maxCompileTime)
	defer cancel()

	// copy files from statically stored Assets to the temp folder
	for name, val := range source {
		ioutil.WriteFile(filepath.Join(tmpDir, name), val, 0644)
	}

	output := bytes.NewBuffer([]byte{})
	cmd := exec.CommandContext(buildCtx, goBin,
		"build",
		"-o",
		"output",
		buildPkgArg)
	cmd.Dir = filepath.Join(tmpDir, "ui")
	cmd.Stderr = output
	cmd.Stdout = output
	goCache := filepath.Join(tmpDir, "gocache")
	cmd.Env = []string{"GOOS=js", "GOARCH=wasm", "GOCACHE=" + goCache, "GOPATH=" + findGopath()}

	err = cmd.Run()
	if err != nil {
		return []byte{}, fmt.Errorf("error compiling: %v\n\nlog output:\n%s", err, output.String())
	}

	res, err := ioutil.ReadFile(filepath.Join(tmpDir, "ui/output"))
	if err != nil {
		return []byte{}, fmt.Errorf("cannot read from file: %v", err)
	}

	return []byte(res), nil
}
