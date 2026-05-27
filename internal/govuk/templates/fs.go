package templates

import (
	"embed"
	"html/template"
	"io/fs"
	"path/filepath"

	"github.com/0xnu/govuk-frontend-go/internal/govuk/templates/engine"
)

//go:embed layouts components pages patterns
var embedded embed.FS

type Set struct {
	T *template.Template
}

func Parse() (*Set, error) {
	t := template.New("root").Funcs(engine.FuncMap())

	root, err := fs.Sub(embedded, ".")
	if err != nil {
		return nil, err
	}

	var patterns []string
	err = fs.WalkDir(root, ".", func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".tmpl" {
			return nil
		}
		patterns = append(patterns, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	_, err = t.ParseFS(root, patterns...)
	if err != nil {
		return nil, err
	}

	return &Set{T: t}, nil
}
