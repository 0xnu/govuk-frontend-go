package contact_a_department_or_service_team

import (
	"html/template"
	"testing"
)

func TestRender(t *testing.T) {
	tpl := template.New("root")
	_, err := tpl.Parse(`{{define "pattern/contact-a-department-or-service-team"}}<div></div>{{end}}`)
	if err != nil {
		t.Fatal(err)
	}
	_, err = Render(tpl, Model{})
	if err != nil {
		t.Fatal(err)
	}
}
