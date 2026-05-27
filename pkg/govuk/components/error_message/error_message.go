package error_message

import (
	"html/template"

	internal "github.com/0xnu/govuk-frontend-go/internal/govuk/components/error_message"
)

type Model = internal.Model

func Render(t *template.Template, m Model) (template.HTML, error) {
	return internal.Render(t, m)
}
