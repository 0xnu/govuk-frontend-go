package summary_list

import (
	"bytes"
	"html/template"
)

func Render(t *template.Template, m Model) (template.HTML, error) {
	var b bytes.Buffer
	err := t.ExecuteTemplate(&b, "govuk/summary-list", m)
	if err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
