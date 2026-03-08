# i18n-timezones-go

> Localized timezone names for Go -- 36 locales, 152 timezones, zero external dependencies.

Building a timezone picker? Displaying meeting times across regions? Go's `time` package gives you IANA identifiers like `America/New_York`, but your users expect to see **"Eastern Time (US & Canada)"** -- in their own language.

**i18n-timezones-go** provides human-friendly, localized timezone display names sourced from [CLDR](https://cldr.unicode.org/), the same data that powers ICU, Chrome, and Android. All translation data is embedded in the binary via `go:embed` -- no external files, no network calls, no filesystem access at runtime.

## Why i18n-timezones-go?

- **36 locales** covering 4+ billion speakers -- from Arabic to Vietnamese
- **Zero external dependencies** -- translation data is embedded at compile time
- **Thread-safe** -- all operations protected by `sync.RWMutex`, safe for concurrent use
- **Register what you need** -- load individual locales or all at once
- **Idiomatic Go** -- returns `(value, bool)` following Go map-access conventions
- **Lightweight** -- only loads locales you register, keeping memory usage minimal

## Install

```bash
go get github.com/onomojo/i18n-timezones-go
```

Requires Go 1.21+.

## Quick Start

```go
package main

import (
    "fmt"
    timezones "github.com/onomojo/i18n-timezones-go"
)

func main() {
    timezones.RegisterLocale("ja")
    timezones.SetDefaultLocale("ja")

    name, _ := timezones.GetTranslation("Tokyo")
    fmt.Println(name) // 東京
}
```

## Usage

### Register only what you need

```go
timezones.RegisterLocale("de")
timezones.RegisterLocale("fr")

name, ok := timezones.GetTranslationForLocale("de", "Tokyo")
// name = "Tokio", ok = true

name, ok = timezones.GetTranslationForLocale("fr", "Eastern Time (US & Canada)")
// name = "Heure de l'Est (États-Unis et Canada)", ok = true
```

### Register all locales at once

For server-side apps where memory is not a concern:

```go
timezones.RegisterAllLocales()
timezones.SetDefaultLocale("en")
```

### Using a default locale

Set a default so you don't have to pass a locale every time:

```go
timezones.RegisterLocale("de")
timezones.SetDefaultLocale("de")

name, _ := timezones.GetTranslation("Tokyo")       // "Tokio"
name, _ = timezones.GetTranslation("Berlin")        // "Berlin"
```

### Get all translations for a locale

```go
all, err := timezones.GetAllTranslations("de")
// all is map[string]string with 152 entries
```

### List available and registered locales

```go
available := timezones.ListLocales()         // all 36 locales in embedded data
registered := timezones.ListRegisteredLocales() // only the ones you've loaded
```

## API Reference

| Function | Description |
|----------|-------------|
| `RegisterLocale(locale string) error` | Load translations for a single locale. No-op if already registered. |
| `RegisterAllLocales() error` | Load translations for all 36 available locales. |
| `SetDefaultLocale(locale string) error` | Set the default locale for lookups. Returns error if locale not registered. |
| `GetDefaultLocale() string` | Get the current default locale, or empty string if none set. |
| `GetTranslation(key string) (string, bool)` | Get the localized name using the default locale. |
| `GetTranslationForLocale(locale, key string) (string, bool)` | Get the localized name for a specific locale. |
| `GetAllTranslations(locale string) (map[string]string, error)` | Get all 152 translations for a locale as a new map. |
| `ListLocales() []string` | List all locales available in the embedded data. |
| `ListRegisteredLocales() []string` | List all currently loaded locales. |

All lookup functions return `(value, false)` when a timezone or locale is not found -- no panics.

## Supported Locales

36 locales covering major world languages:

| | | | | | | |
|---|---|---|---|---|---|---|
| ar | bn | ca | cs | da | de | el |
| en | es | eu | fi | fr | he | hi |
| hr | hu | id | it | ja | ko | ms |
| nl | no | pl | pt | pt-BR | ro | ru |
| sq | sv | th | tr | uk | vi | zh-CN |
| zh-TW | | | | | | |

## Data Source

All translations come from the [Unicode CLDR](https://cldr.unicode.org/) (Common Locale Data Repository) -- the industry-standard source used by every major platform including iOS, Android, Chrome, and Java. This ensures translations are accurate, consistent, and maintained by native speakers through Unicode's established review process.

## Also Available For

- **[Ruby](https://github.com/onomojo/i18n-timezones)** -- Rails gem with automatic `time_zone_select` integration
- **[JavaScript/TypeScript](https://github.com/onomojo/i18n-timezones-js)** -- NPM package with tree-shaking and dropdown helpers
- **[Rust](https://github.com/onomojo/i18n-timezones-rs)** -- Crate with compile-time embedded data

## License

MIT
