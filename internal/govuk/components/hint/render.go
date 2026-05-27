package hint

import (
	"bytes"
	"html/template"
)

func Render(t *template.Template, m Model) (template.HTML, error) {
	var b bytes.Buffer
	err := t.ExecuteTemplate(&b, "govuk/hint", m)
	if err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
