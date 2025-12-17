package hr32

import "fmt"

type ErrGroupIncompleted struct{}

func (e ErrGroupIncompleted) Error() string {
	return "group incompleted"
}

type ErrConfig struct {
	Field       string
	Description string
}

func (e ErrConfig) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Description)
}

type ErrInvalidEncodedLength int

func (e ErrInvalidEncodedLength) Error() string {
	return fmt.Sprintf("invalid encoded length %d", int(e))
}

type ErrInvalidSeparatorIndex int

func (e ErrInvalidSeparatorIndex) Error() string {
	return fmt.Sprintf("invalid separator index %d", int(e))
}

type ErrInvalidChar rune

func (e ErrInvalidChar) Error() string {
	return fmt.Sprintf("invalid character of alphabet: %v", int(e))
}

type ErrNotInitialized struct{}

func (e ErrNotInitialized) Error() string {
	return "config not initialized. Use preset of set all manually with Clone()"
}

type ErrVerify struct{}

func (e ErrVerify) Error() string {
	return "invalid checksum"
}
