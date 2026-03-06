package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestSample2(t *testing.T) {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	sample2()

	w.Close()
	os.Stdout = stdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	expected := "5"
	if len(output) == 0 {
		t.Errorf("sample2() produced no output")
	}
	if !bytes.Contains(buf.Bytes(), []byte(expected)) {
		t.Errorf("sample2() output should contain %q, got %q", expected, output)
	}
}
