## CHANGELOG

### 1.0.6 (21/06/2026)
* Added support for GOV.UK Frontend `6.2.0`
* Refreshed bundled GOV.UK Frontend assets for `6.0.0`, `6.1.0` and `6.2.0` (`current` now points to `6.2.0`)
* Migrated `.golangci.yml` to golangci-lint v2 format so `make lint` works with current tooling

### 1.0.5 (27/05/2026)
* Fixed gofmt -s formatting in internal/app/config/config.go

### 1.0.4 (27/05/2026)
* Fixed CI golangci-lint failure for Go 1.25 by switching GitHub Action install-mode to goinstall

### 1.0.3 (27/05/2026)
* Updated golang.org/x/crypto to fix security alerts
* Updated golang.org/x/net to fix security alerts

### 1.0.2 (27/05/2026)
* Support configuring example listen address via ADDR env var (with PORT fallback)

### 1.0.1 (27/05/2026)
* Fixed invalid default example port (91942 -> 9194)

### 1.0.0 (25/05/2026)
* Initial release
