package timezones

import (
	"testing"
)

func TestRegisterAndTranslate(t *testing.T) {
	// Reset state
	mu.Lock()
	translations = make(map[string]map[string]string)
	defaultLocale = ""
	mu.Unlock()

	if err := RegisterLocale("en"); err != nil {
		t.Fatalf("RegisterLocale(en): %v", err)
	}
	if err := RegisterLocale("ja"); err != nil {
		t.Fatalf("RegisterLocale(ja): %v", err)
	}

	if err := SetDefaultLocale("en"); err != nil {
		t.Fatalf("SetDefaultLocale: %v", err)
	}

	val, ok := GetTranslation("Tokyo")
	if !ok || val != "Tokyo" {
		t.Errorf("GetTranslation(Tokyo) = %q, %v; want Tokyo, true", val, ok)
	}

	val, ok = GetTranslationForLocale("ja", "Tokyo")
	if !ok || val == "" {
		t.Errorf("GetTranslationForLocale(ja, Tokyo) = %q, %v; want non-empty", val, ok)
	}
}

func TestRegisterAllLocales(t *testing.T) {
	mu.Lock()
	translations = make(map[string]map[string]string)
	defaultLocale = ""
	mu.Unlock()

	if err := RegisterAllLocales(); err != nil {
		t.Fatalf("RegisterAllLocales: %v", err)
	}

	locales := ListRegisteredLocales()
	if len(locales) < 36 {
		t.Errorf("expected at least 36 locales, got %d", len(locales))
	}
}

func TestListLocales(t *testing.T) {
	locales := ListLocales()
	if len(locales) < 36 {
		t.Errorf("expected at least 36 available locales, got %d", len(locales))
	}
}

func TestGetAllTranslations(t *testing.T) {
	mu.Lock()
	translations = make(map[string]map[string]string)
	defaultLocale = ""
	mu.Unlock()

	if err := RegisterLocale("en"); err != nil {
		t.Fatalf("RegisterLocale: %v", err)
	}

	all, err := GetAllTranslations("en")
	if err != nil {
		t.Fatalf("GetAllTranslations: %v", err)
	}
	if len(all) < 150 {
		t.Errorf("expected at least 150 translations, got %d", len(all))
	}
}

func TestUnregisteredLocale(t *testing.T) {
	mu.Lock()
	translations = make(map[string]map[string]string)
	defaultLocale = ""
	mu.Unlock()

	err := SetDefaultLocale("zz")
	if err == nil {
		t.Error("SetDefaultLocale(zz) should fail for unregistered locale")
	}

	_, ok := GetTranslationForLocale("zz", "Tokyo")
	if ok {
		t.Error("GetTranslationForLocale(zz) should return false for unregistered locale")
	}
}
