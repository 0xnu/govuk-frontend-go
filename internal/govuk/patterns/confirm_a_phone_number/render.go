package confirm_a_phone_number

import (
	"bytes"
	"html/template"
)

func Render(t *template.Template, m Model) (template.HTML, error) {
	var b bytes.Buffer
	err := t.ExecuteTemplate(&b, "pattern/confirm-a-phone-number", m)
	if err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}
