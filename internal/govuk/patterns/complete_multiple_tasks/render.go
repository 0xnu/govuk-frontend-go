package complete_multiple_tasks

import (
	"bytes"
	"html/template"
)

func Render(t *template.Template, m Model) (template.HTML, error) {
	var b bytes.Buffer
	err := t.ExecuteTemplate(&b, "pattern/complete-multiple-tasks", m)
	if err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
