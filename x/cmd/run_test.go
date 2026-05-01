package cmd

import (
	"testing"
)


func TestIsLocalServer(t *testing.T) {
	tests := []struct {
		name     string
		host     string
		expected bool
	}{
		{
			name:     "empty host (default)",
			host:     "",
			expected: true,
		},
		{
			name:     "localhost",
			host:     "http://localhost:11434",
			expected: true,
		},
		{
			name:     "127.0.0.1",
			host:     "http://127.0.0.1:11434",
			expected: true,
		},
		{
			name:     "custom port on localhost",
			host:     "http://localhost:8080",
			expected: true, // localhost is always considered local
		},
		{
			name:     "remote host",
			host:     "http://ollama.example.com:11434",
			expected: true, // has :11434
		},
		{
			name:     "remote host different port",
			host:     "http://ollama.example.com:8080",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("OLLAMA_HOST", tt.host)
			result := isLocalServer()
			if result != tt.expected {
				t.Errorf("isLocalServer() with OLLAMA_HOST=%q = %v, expected %v", tt.host, result, tt.expected)
			}
		})
	}
}

func TestTruncateToolOutput(t *testing.T) {
	// Create outputs of different sizes
	localLimitOutput := make([]byte, 20000)   // > 4k tokens (16k chars)
	defaultLimitOutput := make([]byte, 50000) // > 10k tokens (40k chars)
	for i := range localLimitOutput {
		localLimitOutput[i] = 'a'
	}
	for i := range defaultLimitOutput {
		defaultLimitOutput[i] = 'b'
	}

	tests := []struct {
		name          string
		output        string
		modelName     string
		host          string
		shouldTrim    bool
		expectedLimit int
	}{
		{
			name:          "short output local model",
			output:        "hello world",
			modelName:     "llama3.2",
			host:          "",
			shouldTrim:    false,
			expectedLimit: localModelTokenLimit,
		},
		{
			name:          "long output local model - trimmed at 4k",
			output:        string(localLimitOutput),
			modelName:     "llama3.2",
			host:          "",
			shouldTrim:    true,
			expectedLimit: localModelTokenLimit,
		},
		{
			name:          "long output remote server - uses 10k limit",
			output:        string(localLimitOutput),
			modelName:     "llama3.2",
			host:          "http://remote.example.com:8080",
			shouldTrim:    false,
			expectedLimit: defaultTokenLimit,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("OLLAMA_HOST", tt.host)
			result := truncateToolOutput(tt.output, tt.modelName)

			if tt.shouldTrim {
				maxLen := tt.expectedLimit * charsPerToken
				if len(result) > maxLen+50 { // +50 for the truncation message
					t.Errorf("expected output to be truncated to ~%d chars, got %d", maxLen, len(result))
				}
				if result == tt.output {
					t.Error("expected output to be truncated but it wasn't")
				}
			} else {
				if result != tt.output {
					t.Error("expected output to not be truncated")
				}
			}
		})
	}
}
