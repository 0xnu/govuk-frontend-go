package assets

import (
	"io/fs"
	"testing"
)

func TestSubdir_CurrentAssets(t *testing.T) {
	sub, err := Subdir("build/govuk-frontend/current")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := fs.Stat(sub, "assets/manifest.json"); err != nil {
		t.Fatal(err)
	}
	if _, err := fs.Stat(sub, "govuk-frontend.min.css"); err != nil {
		t.Fatal(err)
	}
}
