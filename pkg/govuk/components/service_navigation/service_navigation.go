package service_navigation

import (
	"html/template"

	internal "github.com/0xnu/govuk-frontend-go/internal/govuk/components/service_navigation"
)

type Model = internal.Model

func Render(t *template.Template, m Model) (template.HTML, error) {
	return internal.Render(t, m)
}
