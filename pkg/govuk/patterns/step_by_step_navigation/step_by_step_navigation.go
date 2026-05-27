package step_by_step_navigation

import (
	"html/template"

	internal "github.com/0xnu/govuk-frontend-go/internal/govuk/patterns/step_by_step_navigation"
)

type Model = internal.Model

func Render(t *template.Template, m Model) (template.HTML, error) {
	return internal.Render(t, m)
}
