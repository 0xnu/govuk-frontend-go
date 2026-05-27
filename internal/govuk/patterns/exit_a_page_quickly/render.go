package exit_a_page_quickly

import (
	"bytes"
	"html/template"
)

func Render(t *template.Template, m Model) (template.HTML, error) {
	var b bytes.Buffer
	err := t.ExecuteTemplate(&b, "pattern/exit-a-page-quickly", m)
	if err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
