package server

import (
	"github.com/ollama/ollama/internal/modelref"
	"github.com/ollama/ollama/types/model"
)

type modelSource = modelref.ModelSource

const (
	modelSourceUnspecified modelSource = modelref.ModelSourceUnspecified
)

var (
	errModelRequired = modelref.ErrModelRequired
)

type parsedModelRef struct {
	// Original is the caller-provided model string.
	Original string
	// Base is the validated model string.
	Base string
	// Name is Base parsed as a fully-qualified model.Name with defaults applied.
	// Example: "registry.ollama.ai/library/gpt-oss:20b".
	Name model.Name
	// Deprecated in this fork.
	Source modelSource
}

func parseAndValidateModelRef(raw string) (parsedModelRef, error) {
	var zero parsedModelRef

	parsed, err := modelref.ParseRef(raw)
	if err != nil {
		return zero, err
	}

	name := model.ParseName(parsed.Base)
	if !name.IsValid() {
		return zero, model.Unqualified(name)
	}

	return parsedModelRef{
		Original: parsed.Original,
		Base:     parsed.Base,
		Name:     name,
		Source:   parsed.Source,
	}, nil
}

func parseNormalizePullModelRef(raw string) (parsedModelRef, error) { return parseAndValidateModelRef(raw) }
