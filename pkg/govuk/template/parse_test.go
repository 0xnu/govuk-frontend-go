package template

import (
	"bytes"
	"testing"
)

func TestParse(t *testing.T) {
	tpl, err := Parse()
	if err != nil {
		t.Fatal(err)
	}
	if tpl == nil {
		t.Fatal("expected template set")
	}

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, "page", map[string]any{"ServiceName": "Test"}); err != nil {
		t.Fatal(err)
	}
	if buf.Len() == 0 {
		t.Fatal("expected rendered output")
	}
}
