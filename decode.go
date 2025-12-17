package hr32

import (
	"fmt"
	"strings"
)

func (m *Config) Decode(encoded string) (prefix string, version int, data []byte, err error) {
	if !m.initialized {
		return prefix, version, data, ErrNotInitialized{}
	}

	le := len(encoded)

	if le > m.maxLen || le < m.minLen {
		return prefix, version, data, ErrInvalidEncodedLength(le)
	}

	sepLoc := strings.LastIndexByte(encoded, byte(m.separator))
	if sepLoc+7 > le {
		return prefix, version, data, ErrInvalidSeparatorIndex(le)
	}

	prefix = encoded[:sepLoc]
	data, err = m.decode(encoded[sepLoc+1:])

	if err != nil {
		return prefix, version, data, err
	}

	var prefixCond = prefix
	if m.excludePrefix {
		prefixCond = ""
	}

	version = poly(prefixCond, data[:len(data)-6], data[len(data)-6:], m.generators)

	if _, ok := m.versions[version]; !ok {
		return prefix, version, data, ErrVerify{}
	}

	data, err = convert(data[:len(data)-6], 5, 8, false)
	if err != nil {
		return prefix, version, data, fmt.Errorf("can't convert data: %w", err)
	}

	return prefix, version, data, nil
}

func (m *Config) decode(chars string) (decoded []byte, err error) {
	decoded = make([]byte, 0, len(chars))

	for i := 0; i < len(chars); i++ {
		index := strings.IndexByte(m.alphabet, chars[i])

		if index < 0 {
			return nil, ErrInvalidChar(chars[i])
		}

		decoded = append(decoded, byte(index))
	}

	return decoded, nil
}
