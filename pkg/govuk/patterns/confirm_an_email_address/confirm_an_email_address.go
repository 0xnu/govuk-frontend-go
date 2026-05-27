package confirm_an_email_address

import (
	"html/template"

	internal "github.com/0xnu/govuk-frontend-go/internal/govuk/patterns/confirm_an_email_address"
)

type Model = internal.Model

func Render(t *template.Template, m Model) (template.HTML, error) {
	return internal.Render(t, m)
}
