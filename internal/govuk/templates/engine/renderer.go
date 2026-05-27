package engine

import (
	"html/template"
	"io"
	"sync"
)

type Renderer struct {
	mu sync.RWMutex
	t  *template.Template
}

func New(t *template.Template) *Renderer {
	return &Renderer{t: t}
}

func (r *Renderer) Execute(w io.Writer, name string, data any) error {
	r.mu.RLock()
	t := r.t
	r.mu.RUnlock()
	return t.ExecuteTemplate(w, name, data)
}
