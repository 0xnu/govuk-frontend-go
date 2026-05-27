package confirm_a_phone_number

import (
	"html/template"

	internal "github.com/0xnu/govuk-frontend-go/internal/govuk/patterns/confirm_a_phone_number"
)

type Model = internal.Model

func Render(t *template.Template, m Model) (template.HTML, error) {
	return internal.Render(t, m)
}
