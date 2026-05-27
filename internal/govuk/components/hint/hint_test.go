package hint

import (
	"html/template"
	"testing"
)

func TestRender(t *testing.T) {
	tpl := template.New("root")
	_, err := tpl.Parse(`{{define "govuk/hint"}}<div>{{.Text}}</div>{{end}}`)
	if err != nil {
		t.Fatal(err)
	}
	out, err := Render(tpl, Model{Text: "Help text"})
	if err != nil {
		t.Fatal(err)
	}
	if string(out) == "" {
		t.Fatal("expected html")
	}
}
