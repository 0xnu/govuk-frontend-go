package label

import (
	"html/template"
	"testing"
)

func TestRender(t *testing.T) {
	tpl := template.New("root")
	_, err := tpl.Parse(`{{define "govuk/label"}}<label>{{.Text}}</label>{{end}}`)
	if err != nil {
		t.Fatal(err)
	}
	out, err := Render(tpl, Model{Text: "Name"})
	if err != nil {
		t.Fatal(err)
	}
	if string(out) == "" {
		t.Fatal("expected html")
	}
}
