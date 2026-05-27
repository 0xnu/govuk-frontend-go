package confirm_an_email_address

import (
	"bytes"
	"html/template"
)

func Render(t *template.Template, m Model) (template.HTML, error) {
	var b bytes.Buffer
	err := t.ExecuteTemplate(&b, "pattern/confirm-an-email-address", m)
	if err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
