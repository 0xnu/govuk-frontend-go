package national_insurance_numbers

import (
	"html/template"
	"testing"
)

func TestRender(t *testing.T) {
	tpl := template.New("root")
	_, err := tpl.Parse(`{{define "pattern/national-insurance-numbers"}}<div></div>{{end}}`)
	if err != nil {
		t.Fatal(err)
	}
	_, err = Render(tpl, Model{})
	if err != nil {
		t.Fatal(err)
	}
}
