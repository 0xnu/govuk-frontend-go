package breadcrumbs

import (
	"html/template"

	internal "github.com/0xnu/govuk-frontend-go/internal/govuk/components/breadcrumbs"
)

type Model = internal.Model

func Render(t *template.Template, m Model) (template.HTML, error) {
	return internal.Render(t, m)
}
