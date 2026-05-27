package config

import "testing"

func TestLoad_Default(t *testing.T) {
	t.Setenv("PORT", "")
	c := Load()
	if c.Port != "91942" {
		t.Fatalf("expected default port 91942, got %q", c.Port)
	}
	if c.Addr() != ":91942" {
		t.Fatalf("expected addr :91942, got %q", c.Addr())
	}
}

func TestLoad_FromEnv(t *testing.T) {
	t.Setenv("PORT", "1234")
	c := Load()
	if c.Port != "1234" {
		t.Fatalf("expected port 1234, got %q", c.Port)
	}
	if c.Addr() != ":1234" {
		t.Fatalf("expected addr :1234, got %q", c.Addr())
	}
}
