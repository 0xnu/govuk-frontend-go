package task_list

import (
	"html/template"

	internal "github.com/0xnu/govuk-frontend-go/internal/govuk/components/task_list"
)

type Model = internal.Model

func Render(t *template.Template, m Model) (template.HTML, error) {
	return internal.Render(t, m)
}
