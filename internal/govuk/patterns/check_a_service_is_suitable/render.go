package check_a_service_is_suitable

import (
	"bytes"
	"html/template"
)

func Render(t *template.Template, m Model) (template.HTML, error) {
	var b bytes.Buffer
	err := t.ExecuteTemplate(&b, "pattern/check-a-service-is-suitable", m)
	if err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
