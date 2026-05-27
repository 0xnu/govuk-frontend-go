package button

import (
	"html/template"
	"testing"
)

func TestRender(t *testing.T) {
	tpl := template.New("root")
	_, err := tpl.Parse(`{{define "govuk/button"}}<button>{{.Text}}</button>{{end}}`)
	if err != nil {
		t.Fatal(err)
	}
	out, err := Render(tpl, Model{Text: "Continue"})
	if err != nil {
		t.Fatal(err)
	}
	if string(out) == "" {
		t.Fatal("expected html")
	}
}
