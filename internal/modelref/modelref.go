package modelref

import (
	"errors"
	"strings"
)

type ModelSource uint8

const (
	ModelSourceUnspecified ModelSource = iota
)

var (
	ErrModelRequired = errors.New("model is required")
)

type ParsedRef struct {
	Original string
	Base     string
	Source   ModelSource
}

func ParseRef(raw string) (ParsedRef, error) {
	var zero ParsedRef

	raw = strings.TrimSpace(raw)
	if raw == "" {
		return zero, ErrModelRequired
	}

	return ParsedRef{
		Original: raw,
		Base:     raw,
		Source:   ModelSourceUnspecified,
	}, nil
}
