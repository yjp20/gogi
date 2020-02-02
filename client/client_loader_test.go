package client

import (
	"testing"
)

func TestCompileIndex(t *testing.T) {
	_, err := CompileIndex(&struct{}{})
	if err != nil {
		t.Error(err)
	}
}
