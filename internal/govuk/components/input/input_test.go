package input

import (
	"html/template"
	"testing"
)

func TestRender(t *testing.T) {
	tpl := template.New("root")
	_, err := tpl.Parse(`{{define "govuk/input"}}<input id="{{.ID}}">{{end}}`)
	if err != nil {
		t.Fatal(err)
	}
	out, err := Render(tpl, Model{ID: "name"})
	if err != nil {
		t.Fatal(err)
	}
	if string(out) == "" {
		t.Fatal("expected html")
	}
}
