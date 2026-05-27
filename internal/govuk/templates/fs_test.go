package templates

import "testing"

func TestParseIncludesExamplePages(t *testing.T) {
	set, err := Parse()
	if err != nil {
		t.Fatal(err)
	}
	if set.T.Lookup("page") == nil {
		t.Fatal(`expected template "page"`)
	}
	if set.T.Lookup("form") == nil {
		t.Fatal(`expected template "form"`)
	}
}
