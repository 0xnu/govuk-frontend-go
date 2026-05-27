package assets

import "testing"

func TestDefaultPrefix(t *testing.T) {
	if DefaultPrefix != "/assets" {
		t.Fatalf("expected DefaultPrefix %q, got %q", "/assets", DefaultPrefix)
	}
}
