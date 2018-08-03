package llog

import (
	"fmt"
	"strings"
)

// Level type
type Level int32

// Convert the Level to a string. E.g. PanicLevel becomes "panic".
func (level Level) String() string {
	switch level {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	}

	return "unknown"
}

// ParseLevel takes a string level and returns the level enum
func ParseLevel(lvl string) (Level, error) {
	switch strings.ToLower(lvl) {
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	}

	var l Level
	return l, fmt.Errorf("not a valid Level: %q", lvl)
}

// AllLevels is a constant exposing all logging levels
var AllLevels = []Level{
	InfoLevel,
	DebugLevel,
}

// enums for error levels
const (
	DebugLevel Level = iota
	InfoLevel
)
