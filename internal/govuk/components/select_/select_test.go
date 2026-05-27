package select_

import (
	"html/template"
	"testing"
)

func TestRender(t *testing.T) {
	tpl := template.New("root")
	_, err := tpl.Parse(`{{define "govuk/select"}}<div></div>{{end}}`)
	if err != nil {
		t.Fatal(err)
	}
	_, err = Render(tpl, Model{})
	if err != nil {
		t.Fatal(err)
	}
}
