package gin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSetGovUKFrontendVersion(t *testing.T) {
	a := New()

	if err := a.SetGovUKFrontendVersion("6.1.0"); err != nil {
		t.Fatal(err)
	}
	if a.AssetsRoot != "build/govuk-frontend/v6.1.0" {
		t.Fatalf("unexpected AssetsRoot: %q", a.AssetsRoot)
	}

	if err := a.SetGovUKFrontendVersion("does-not-exist"); err == nil {
		t.Fatal("expected error")
	}
}

func TestMount_ServesAssetsAndFavicon(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	a := New()
	a.Mount(r)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/assets/govuk-frontend.min.css", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
	if w.Body.Len() == 0 {
		t.Fatal("expected non-empty body")
	}

	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/assets/does-not-exist", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", w.Code)
	}

	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/favicon.ico", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}
