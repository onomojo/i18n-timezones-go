# i18n-timezones-go

Go module for localized timezone name translations. Covers 152 ActiveSupport timezones across 36 locales, sourced from Unicode CLDR.

## Install

```bash
go get github.com/onomojo/i18n-timezones-go
```

## Usage

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

## API

- `RegisterLocale(locale string) error`
- `RegisterAllLocales() error`
- `SetDefaultLocale(locale string) error`
- `GetDefaultLocale() string`
- `GetTranslation(key string) (string, bool)`
- `GetTranslationForLocale(locale, key string) (string, bool)`
- `GetAllTranslations(locale string) (map[string]string, error)`
- `ListLocales() []string` — available in embedded data
- `ListRegisteredLocales() []string` — currently loaded

## License

MIT
