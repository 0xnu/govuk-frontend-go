package error_summary

import (
	"bytes"
	"html/template"
)

func Render(t *template.Template, m Model) (template.HTML, error) {
	var b bytes.Buffer
	err := t.ExecuteTemplate(&b, "govuk/error-summary", m)
	if err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
