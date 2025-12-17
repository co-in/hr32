package hr32

import (
	"fmt"
	"strings"
)

// Len counts the length without actual encoding.
// It returns the overflow=true parameter if it is greater than the config limit.
// You can also override the config parameters through options
func (m *Config) Len(data []byte, options ...Option) (length int, err error) {
	if !m.initialized {
		return length, ErrNotInitialized{}
	}

	cfg := *m

	for _, option := range options {
		option(&cfg)
	}

	length = len(cfg.prefix) + 1 + (len(data)*8+4)/5 + 6

	if length > cfg.maxLen {
		return length, ErrInvalidEncodedLength(length)
	}

	return length, nil
}

// Encode returns an unserialized structure, with all necessary fields that can be processed.
// For example, to format your additional formatting, which can be processed back before Decode
func (m *Config) Encode(data []byte) (result string, err error) {
	if !m.initialized {
		return result, ErrNotInitialized{}
	}

	var converted []byte

	if converted, err = convert(data, 8, 5, true); err != nil {
		return result, fmt.Errorf("can't convert data: %w", err)
	}

	var prefix = m.prefix

	if m.excludePrefix {
		prefix = ""
	}

	var checksum [6]byte
	version := poly(prefix, converted, nil, m.generators) ^ m.version

	calcChecksum(&checksum, version, m.alphabet)

	for i, chr := range converted {
		converted[i] = m.alphabet[chr]
	}

	var builder strings.Builder
	builder.Grow(len(m.prefix) + 1 + len(converted) + 6)

	if m.prefix != "" {
		builder.WriteString(m.prefix)
		builder.WriteRune(m.separator)
	}

	builder.Write(converted)
	builder.Write(checksum[:])

	return builder.String(), nil
}

func (m *Config) Convert(encoded string, callback func(prefix string, version int) *Config) (result string, err error) {
	prefix, version, data, err := m.Decode(encoded)
	if err != nil {
		return result, err
	}

	return callback(prefix, version).Encode(data)
}
