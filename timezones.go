package timezones

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"

	data "github.com/onomojo/i18n-timezones-data"
)

var (
	mu            sync.RWMutex
	translations  = make(map[string]map[string]string)
	defaultLocale string
)

// RegisterLocale loads translations for the given locale from embedded data.
func RegisterLocale(locale string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := translations[locale]; ok {
		return nil
	}

	m, err := loadLocale(locale)
	if err != nil {
		return err
	}
	translations[locale] = m
	return nil
}

// RegisterAllLocales loads translations for every available locale.
func RegisterAllLocales() error {
	locales, err := listAvailableLocales()
	if err != nil {
		return err
	}
	mu.Lock()
	defer mu.Unlock()

	for _, locale := range locales {
		if _, ok := translations[locale]; ok {
			continue
		}
		m, err := loadLocale(locale)
		if err != nil {
			return err
		}
		translations[locale] = m
	}
	return nil
}

// SetDefaultLocale sets the default locale for lookups.
func SetDefaultLocale(locale string) error {
	mu.RLock()
	_, ok := translations[locale]
	mu.RUnlock()
	if !ok {
		return fmt.Errorf("timezones: locale %q not registered", locale)
	}
	mu.Lock()
	defaultLocale = locale
	mu.Unlock()
	return nil
}

// GetDefaultLocale returns the current default locale.
func GetDefaultLocale() string {
	mu.RLock()
	defer mu.RUnlock()
	return defaultLocale
}

// GetTranslation returns the translation for a timezone key using the default locale.
func GetTranslation(key string) (string, bool) {
	mu.RLock()
	defer mu.RUnlock()
	if defaultLocale == "" {
		return "", false
	}
	m, ok := translations[defaultLocale]
	if !ok {
		return "", false
	}
	v, ok := m[key]
	return v, ok
}

// GetTranslationForLocale returns the translation for a timezone key in a specific locale.
func GetTranslationForLocale(locale, key string) (string, bool) {
	mu.RLock()
	defer mu.RUnlock()
	m, ok := translations[locale]
	if !ok {
		return "", false
	}
	v, ok := m[key]
	return v, ok
}

// GetAllTranslations returns all translations for a locale.
func GetAllTranslations(locale string) (map[string]string, error) {
	mu.RLock()
	defer mu.RUnlock()
	m, ok := translations[locale]
	if !ok {
		return nil, fmt.Errorf("timezones: locale %q not registered", locale)
	}
	result := make(map[string]string, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result, nil
}

// ListLocales returns all locales available in the embedded data.
func ListLocales() []string {
	locales, err := listAvailableLocales()
	if err != nil {
		return nil
	}
	return locales
}

// ListRegisteredLocales returns all currently loaded locales.
func ListRegisteredLocales() []string {
	mu.RLock()
	defer mu.RUnlock()
	result := make([]string, 0, len(translations))
	for k := range translations {
		result = append(result, k)
	}
	sort.Strings(result)
	return result
}

func loadLocale(locale string) (map[string]string, error) {
	b, err := data.FS.ReadFile("data/" + locale + ".json")
	if err != nil {
		return nil, fmt.Errorf("timezones: locale %q not found: %w", locale, err)
	}
	var m map[string]string
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, fmt.Errorf("timezones: failed to parse %q: %w", locale, err)
	}
	return m, nil
}

func listAvailableLocales() ([]string, error) {
	entries, err := data.FS.ReadDir("data")
	if err != nil {
		return nil, err
	}
	var locales []string
	for _, e := range entries {
		name := e.Name()
		if strings.HasSuffix(name, ".json") {
			locales = append(locales, strings.TrimSuffix(name, ".json"))
		}
	}
	sort.Strings(locales)
	return locales, nil
}
