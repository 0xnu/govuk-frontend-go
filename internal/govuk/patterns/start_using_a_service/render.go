package start_using_a_service

import (
	"bytes"
	"html/template"
)

func Render(t *template.Template, m Model) (template.HTML, error) {
	var b bytes.Buffer
	err := t.ExecuteTemplate(&b, "pattern/start-using-a-service", m)
	if err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
