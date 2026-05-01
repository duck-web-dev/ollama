package modelref

import (
	"errors"
	"testing"
)

func TestParseRef(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantBase   string
		wantSource ModelSource
		wantErr    error
	}{
		{name: "plain", input: "llama3.2", wantBase: "llama3.2", wantSource: ModelSourceUnspecified},
		{name: "tagged", input: "gpt-oss:20b", wantBase: "gpt-oss:20b", wantSource: ModelSourceUnspecified},
		{name: "trimmed", input: "  qwen3:8b  ", wantBase: "qwen3:8b", wantSource: ModelSourceUnspecified},
		{name: "empty", input: "   ", wantErr: ErrModelRequired, wantSource: ModelSourceUnspecified},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRef(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("ParseRef(%q) error = %v, want %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseRef(%q) returned error: %v", tt.input, err)
			}

			if got.Base != tt.wantBase {
				t.Fatalf("base = %q, want %q", got.Base, tt.wantBase)
			}

			if got.Source != tt.wantSource {
				t.Fatalf("source = %v, want %v", got.Source, tt.wantSource)
			}
		})
	}
}
