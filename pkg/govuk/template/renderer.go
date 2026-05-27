package template

import (
	"html/template"
	"io"
)

type Renderer interface {
	Execute(w io.Writer, name string, data any) error
}

type Templates struct {
	T *template.Template
}

func (t Templates) Execute(w io.Writer, name string, data any) error {
	return t.T.ExecuteTemplate(w, name, data)
}
