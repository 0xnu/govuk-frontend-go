package validation

import "testing"

func TestErrors_AddAndQuery(t *testing.T) {
	var errs Errors

	errs.Add("name", "Required")
	errs.Add("email", "Invalid")
	errs.AddGlobal("Something went wrong")

	if !errs.Any() {
		t.Fatalf("expected Any() to be true")
	}

	if !errs.HasField("name") {
		t.Fatalf("expected HasField(name) to be true")
	}

	if errs.HasField("missing") {
		t.Fatalf("expected HasField(missing) to be false")
	}

	nameErrs := errs.ForField("name")
	if len(nameErrs) != 1 {
		t.Fatalf("expected 1 error for name, got %d", len(nameErrs))
	}
	if nameErrs[0].Message != "Required" {
		t.Fatalf("expected message %q, got %q", "Required", nameErrs[0].Message)
	}

	first, ok := errs.First("email")
	if !ok {
		t.Fatalf("expected First(email) ok=true")
	}
	if first.Message != "Invalid" {
		t.Fatalf("expected message %q, got %q", "Invalid", first.Message)
	}

	_, ok = errs.First("missing")
	if ok {
		t.Fatalf("expected First(missing) ok=false")
	}
}
