package confirm_a_phone_number

import (
	"html/template"
	"testing"
)

func TestRender(t *testing.T) {
	tpl := template.New("root")
	_, err := tpl.Parse(`{{define "pattern/confirm-a-phone-number"}}<div></div>{{end}}`)
	if err != nil {
		t.Fatal(err)
	}
	_, err = Render(tpl, Model{})
	if err != nil {
		t.Fatal(err)
	}
}
