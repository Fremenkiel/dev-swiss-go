package handlers

import (
	"bytes"
	"io"
	"os"
	"sync"
	"testing"
)

func setupEnv(t *testing.T, args []string) (*os.File, *os.File, *sync.WaitGroup, *bytes.Buffer) {
	os.Args = args

	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("Unable to create pipe")
	}

	var buf bytes.Buffer
	var wg sync.WaitGroup

	wg.Go(func() {
		io.Copy(&buf, r)
	})

	os.Stdout = w

	return r, w, &wg, &buf
}

func cleanupEnv(args []string, stdout *os.File) {
	os.Args = args
	os.Stdout = stdout
}
