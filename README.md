## govuk-frontend-go

[![Release](https://img.shields.io/github/release/0xnu/govuk-frontend-go.svg)](https://github.com/0xnu/govuk-frontend-go/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/0xnu/govuk-frontend-go)](https://goreportcard.com/report/github.com/0xnu/govuk-frontend-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/0xnu/govuk-frontend-go.svg)](https://pkg.go.dev/github.com/0xnu/govuk-frontend-go)
[![License](https://img.shields.io/github/license/0xnu/govuk-frontend-go)](/LICENSE)

GOV.UK Design System integration for [Go](https://go.dev/), with a [Gin](https://gin-gonic.com/en/) adapter.

It provides:
- Versioned GOV.UK Frontend assets
- Go template helpers and component renderers
- Gin integration for serving assets and rendering templates
- [A coverage matrix](/DESIGN_SYSTEM_COVERAGE.md)

### Use In Your Gin App (Recommended)

Add the module to your service:

```bash
go get github.com/0xnu/govuk-frontend-go@<version-tag>
```

Minimal setup:

```go
package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	govukgin "github.com/0xnu/govuk-frontend-go/pkg/govuk/gin"
	govuktemplate "github.com/0xnu/govuk-frontend-go/pkg/govuk/template"
)

func main() {
	r := gin.Default()

	t, err := govuktemplate.Parse()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)

	g := govukgin.New()
	if err := g.SetGovUKFrontendVersion("6.1.0"); err != nil {
		panic(err)
	}
	g.Mount(r)

	r.GET("/", handleIndex)

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":91942"
	}

	if err := r.Run(addr); err != nil {
		panic(err)
	}
}

func handleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "page", gin.H{
		"ServiceName": "Example service",
	})
}
```

What this gives you:
- Templates embedded and loaded via `pkg/govuk/template` (source templates live under [internal/govuk/templates](/internal/govuk/templates))
- Embedded assets mounted at:
  - `GET /assets/*`
  - `GET /favicon.ico`
 
Coverage:
- Components and patterns implemented in this module are listed in [DESIGN_SYSTEM_COVERAGE.md](/DESIGN_SYSTEM_COVERAGE.md).

Versioning:
- Prefer a tag (e.g. `@v1.0.0`) rather than `@main`.
- Embedded GOV.UK Frontend assets are shipped for: `6.0.0`, `6.1.0` (and `current` defaults to the latest bundled version).

Select the GOV.UK Frontend version you want to serve:

```go
g := govukgin.New()
if err := g.SetGovUKFrontendVersion("6.1.0"); err != nil {
	panic(err)
}
g.Mount(r)
```

### Quickstart

```bash
go run ./cmd/example-gin-app
```

- Home page: `http://localhost:91942/`

### License

This project is licensed under the [MIT License](./LICENSE).

### Copyright

(c) 2026 [Finbarrs Oketunji](https://finbarrs.eu). All Rights Reserved.

govuk-frontend-go is Licensed under the [Open Government Licence v3.0](https://www.nationalarchives.gov.uk/doc/open-government-licence/version/3/)
