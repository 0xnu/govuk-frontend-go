package error_summary

import (
	"html/template"
	"testing"
)

func TestRender(t *testing.T) {
	tpl := template.New("root")
	_, err := tpl.Parse(`{{define "govuk/error-summary"}}<div></div>{{end}}`)
	if err != nil {
		t.Fatal(err)
	}
	out, err := Render(tpl, Model{})
	if err != nil {
		t.Fatal(err)
	}
	if string(out) == "" {
		t.Fatal("expected html")
	}

	_, err = Render(tpl, Model{Title: "There is a problem", Items: []Item{{Href: "#name", Text: "Enter a name"}}})
	if err != nil {
		t.Fatal(err)
	}
}
