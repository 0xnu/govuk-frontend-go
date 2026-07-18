package generic_header

import (
	"html/template"
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	tpl := template.New("root")
	_, err := tpl.Parse(`{{define "govuk/generic-header"}}<div class="govuk-generic-header"><a href="{{.URL}}">{{.LogoText}}</a></div>{{end}}`)
	if err != nil {
		t.Fatal(err)
	}
	html, err := Render(tpl, Model{URL: "/", LogoText: "Test"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(html), "Test") {
		t.Fatalf("unexpected output: %s", html)
	}
}
