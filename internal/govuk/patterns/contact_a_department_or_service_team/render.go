package contact_a_department_or_service_team

import (
	"bytes"
	"html/template"
)

func Render(t *template.Template, m Model) (template.HTML, error) {
	var b bytes.Buffer
	err := t.ExecuteTemplate(&b, "pattern/contact-a-department-or-service-team", m)
	if err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
