package phone_numbers

import (
	"html/template"

	internal "github.com/0xnu/govuk-frontend-go/internal/govuk/patterns/phone_numbers"
)

type Model = internal.Model

func Render(t *template.Template, m Model) (template.HTML, error) {
	return internal.Render(t, m)
}
