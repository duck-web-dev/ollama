package server

import (
	"errors"
	"strings"
	"testing"
)

func TestParseModelRef(t *testing.T) {
	t.Run("valid model", func(t *testing.T) {
		got, err := parseAndValidateModelRef("gpt-oss:20b")
		if err != nil {
			t.Fatalf("parseAndValidateModelRef returned error: %v", err)
		}

		if got.Base != "gpt-oss:20b" {
			t.Fatalf("expected base gpt-oss:20b, got %q", got.Base)
		}

		if got.Name.String() != "registry.ollama.ai/library/gpt-oss:20b" {
			t.Fatalf("unexpected resolved name: %q", got.Name.String())
		}
	})

	t.Run("defaults tag", func(t *testing.T) {
		got, err := parseAndValidateModelRef("llama3")
		if err != nil {
			t.Fatalf("parseAndValidateModelRef returned error: %v", err)
		}

		if got.Name.Tag != "latest" {
			t.Fatalf("expected default latest tag, got %q", got.Name.Tag)
		}
	})

	t.Run("empty model fails", func(t *testing.T) {
		_, err := parseAndValidateModelRef("")
		if !errors.Is(err, errModelRequired) {
			t.Fatalf("expected errModelRequired, got %v", err)
		}
	})

	t.Run("invalid model fails", func(t *testing.T) {
		_, err := parseAndValidateModelRef("::bad")
		if err == nil {
			t.Fatal("expected error for invalid model")
		}
		if !strings.Contains(err.Error(), "unqualified") {
			t.Fatalf("expected unqualified model error, got %v", err)
		}
	})
}

func TestParsePullModelRef(t *testing.T) {
	got, err := parseNormalizePullModelRef("gpt-oss:20b")
	if err != nil {
		t.Fatalf("parseNormalizePullModelRef returned error: %v", err)
	}

	if got.Base != "gpt-oss:20b" {
		t.Fatalf("expected base gpt-oss:20b, got %q", got.Base)
	}

	if got.Name.String() != "registry.ollama.ai/library/gpt-oss:20b" {
		t.Fatalf("unexpected resolved name: %q", got.Name.String())
	}
}
