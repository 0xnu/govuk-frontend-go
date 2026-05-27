package file_upload

import (
	"bytes"
	"html/template"
)

func Render(t *template.Template, m Model) (template.HTML, error) {
	var b bytes.Buffer
	err := t.ExecuteTemplate(&b, "govuk/file-upload", m)
	if err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
