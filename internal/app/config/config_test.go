package config

import "testing"

func TestLoad_Default(t *testing.T) {
	t.Setenv("ADDR", "")
	t.Setenv("PORT", "")
	c := Load()
	if c.Port != "9194" {
		t.Fatalf("expected default port 9194, got %q", c.Port)
	}
	if c.Addr() != ":9194" {
		t.Fatalf("expected addr :9194, got %q", c.Addr())
	}
}

func TestLoad_FromEnv(t *testing.T) {
	t.Setenv("ADDR", "")
	t.Setenv("PORT", "1234")
	c := Load()
	if c.Port != "1234" {
		t.Fatalf("expected port 1234, got %q", c.Port)
	}
	if c.Addr() != ":1234" {
		t.Fatalf("expected addr :1234, got %q", c.Addr())
	}
}

func TestLoad_FromAddrEnv(t *testing.T) {
	t.Setenv("ADDR", "0.0.0.0:5555")
	t.Setenv("PORT", "")
	c := Load()
	if c.Port != "5555" {
		t.Fatalf("expected port 5555, got %q", c.Port)
	}
	if c.Addr() != "0.0.0.0:5555" {
		t.Fatalf("expected addr 0.0.0.0:5555, got %q", c.Addr())
	}
}
