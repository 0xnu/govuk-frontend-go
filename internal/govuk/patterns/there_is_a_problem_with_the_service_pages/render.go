package there_is_a_problem_with_the_service_pages

import (
	"bytes"
	"html/template"
)

func Render(t *template.Template, m Model) (template.HTML, error) {
	var b bytes.Buffer
	err := t.ExecuteTemplate(&b, "pattern/there-is-a-problem-with-the-service-pages", m)
	if err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
