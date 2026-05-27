package recover_from_validation_errors

import (
	"bytes"
	"html/template"
)

func Render(t *template.Template, m Model) (template.HTML, error) {
	var b bytes.Buffer
	err := t.ExecuteTemplate(&b, "pattern/recover-from-validation-errors", m)
	if err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
