package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestSample(t *testing.T) {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	sample()

	w.Close()
	os.Stdout = stdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	expected := "4"
	if len(output) == 0 {
		t.Errorf("sample() produced no output")
	}
	if !bytes.Contains(buf.Bytes(), []byte(expected)) {
		t.Errorf("sample() output should contain %q, got %q", expected, output)
	}
}
