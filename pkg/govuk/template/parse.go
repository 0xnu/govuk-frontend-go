package template

import (
	"html/template"

	internal "github.com/0xnu/govuk-frontend-go/internal/govuk/templates"
)

func Parse() (*template.Template, error) {
	set, err := internal.Parse()
	if err != nil {
		return nil, err
	}
	return set.T, nil
}
