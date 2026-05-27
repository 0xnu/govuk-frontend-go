package api

import "testing"

func TestTypes(t *testing.T) {
	var a Attrs = map[string]string{"k": "v"}
	if a["k"] != "v" {
		t.Fatalf("expected Attrs to behave like map, got %q", a["k"])
	}

	var h HTML = "<p>x</p>"
	if string(h) != "<p>x</p>" {
		t.Fatalf("unexpected HTML value: %q", h)
	}
}
