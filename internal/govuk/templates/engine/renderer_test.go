package engine

import (
	"bytes"
	"html/template"
	"testing"
)

func TestRenderer_Execute(t *testing.T) {
	tpl := template.New("root")
	_, err := tpl.Parse(`{{define "x"}}Hello {{.}}{{end}}`)
	if err != nil {
		t.Fatal(err)
	}

	r := New(tpl)

	var buf bytes.Buffer
	if err := r.Execute(&buf, "x", "world"); err != nil {
		t.Fatal(err)
	}
	if buf.String() != "Hello world" {
		t.Fatalf("unexpected output: %q", buf.String())
	}
}
