package create_a_username

import (
	"html/template"

	internal "github.com/0xnu/govuk-frontend-go/internal/govuk/patterns/create_a_username"
)

type Model = internal.Model

func Render(t *template.Template, m Model) (template.HTML, error) {
	return internal.Render(t, m)
}
